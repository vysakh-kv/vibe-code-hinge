package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// ProfileService handles profile-related business logic
type ProfileService struct {
	db *sql.DB
}

// NewProfileService creates a new profile service
func NewProfileService(db *sql.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}

// GetProfileByID retrieves a profile by ID
func (s *ProfileService) GetProfileByID(ctx context.Context, id int64) (*models.Profile, error) {
	// Get profile
	var profile models.Profile
	err := s.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, name, bio, date_of_birth, gender, location, occupation, created_at, updated_at 
		FROM profiles WHERE id = $1`,
		id,
	).Scan(
		&profile.ID, &profile.UserID, &profile.Name, &profile.Bio, &profile.DateOfBirth,
		&profile.Gender, &profile.Location, &profile.Occupation, &profile.CreatedAt, &profile.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	// Get photos
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT id, profile_id, url, is_primary, created_at 
		FROM photos WHERE profile_id = $1 ORDER BY is_primary DESC, created_at ASC`,
		profile.ID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []models.Photo
	for rows.Next() {
		var photo models.Photo
		if err := rows.Scan(&photo.ID, &photo.ProfileID, &photo.URL, &photo.IsPrimary, &photo.CreatedAt); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	profile.Photos = photos
	return &profile, nil
}

// CreateOrUpdateProfile creates or updates a profile
func (s *ProfileService) CreateOrUpdateProfile(ctx context.Context, userID int64, input models.ProfileInput) (*models.Profile, error) {
	// Parse date of birth
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		return nil, errors.New("invalid date format for date of birth")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check if profile exists
	var profileID int64
	var exists bool
	err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM profiles WHERE user_id = $1)",
		userID,
	).Scan(&exists)

	if err != nil {
		return nil, err
	}

	now := time.Now()
	if exists {
		// Update existing profile
		err = tx.QueryRowContext(
			ctx,
			`UPDATE profiles SET name = $1, bio = $2, date_of_birth = $3, gender = $4, location = $5, occupation = $6, updated_at = $7
			WHERE user_id = $8 RETURNING id`,
			input.Name, input.Bio, dob, input.Gender, input.Location, input.Occupation, now, userID,
		).Scan(&profileID)
	} else {
		// Create new profile
		err = tx.QueryRowContext(
			ctx,
			`INSERT INTO profiles (user_id, name, bio, date_of_birth, gender, location, occupation, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
			userID, input.Name, input.Bio, dob, input.Gender, input.Location, input.Occupation, now, now,
		).Scan(&profileID)
	}

	if err != nil {
		return nil, err
	}

	// Process photos if provided
	if len(input.Photos) > 0 {
		// Create new photos
		for i, photoURL := range input.Photos {
			isPrimary := i == 0 // First photo is primary
			_, err := tx.ExecContext(
				ctx,
				`INSERT INTO photos (profile_id, url, is_primary, created_at)
				VALUES ($1, $2, $3, $4)`,
				profileID, photoURL, isPrimary, now,
			)
			if err != nil {
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Get updated profile
	return s.GetProfileByID(ctx, profileID)
}
