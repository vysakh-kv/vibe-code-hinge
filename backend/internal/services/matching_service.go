package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// MatchingService handles matching-related business logic
type MatchingService struct {
	db *sql.DB
}

// NewMatchingService creates a new matching service
func NewMatchingService(db *sql.DB) *MatchingService {
	return &MatchingService{
		db: db,
	}
}

// GetDiscoverProfiles retrieves profiles for discovery based on user preferences
func (s *MatchingService) GetDiscoverProfiles(ctx context.Context, userID int64, limit int) ([]*models.Profile, error) {
	// Get profiles that have not been swiped by the user
	// Implement filtering by preferences later
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT p.id, p.user_id, p.name, p.bio, p.date_of_birth, p.gender, p.location, p.occupation, p.created_at, p.updated_at
		FROM profiles p
		WHERE p.user_id != $1 AND p.id NOT IN (
			SELECT profile_id FROM swipes WHERE user_id = $1
		)
		LIMIT $2`,
		userID, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		err := rows.Scan(
			&profile.ID, &profile.UserID, &profile.Name, &profile.Bio, &profile.DateOfBirth,
			&profile.Gender, &profile.Location, &profile.Occupation, &profile.CreatedAt, &profile.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Get photos for each profile
		profile.Photos, err = s.getProfilePhotos(ctx, profile.ID)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, &profile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

// CreateSwipe records a user's swipe on a profile
func (s *MatchingService) CreateSwipe(ctx context.Context, userID, profileID int64, isLike bool) (*models.Match, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Record the swipe
	now := time.Now()
	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO swipes (user_id, profile_id, is_like, created_at)
		VALUES ($1, $2, $3, $4)`,
		userID, profileID, isLike, now,
	)
	if err != nil {
		return nil, err
	}

	// If this was a like, check if there's a mutual like
	var match *models.Match
	if isLike {
		// Get the user ID of the profile owner
		var otherUserID int64
		err = tx.QueryRowContext(
			ctx,
			"SELECT user_id FROM profiles WHERE id = $1",
			profileID,
		).Scan(&otherUserID)
		if err != nil {
			return nil, err
		}

		// Check if the other user has liked this user
		var mutualLike bool
		err = tx.QueryRowContext(
			ctx,
			`SELECT EXISTS(
				SELECT 1 FROM swipes s
				JOIN profiles p ON s.profile_id = p.id
				WHERE s.user_id = $1 AND p.user_id = $2 AND s.is_like = true
			)`,
			otherUserID, userID,
		).Scan(&mutualLike)
		if err != nil {
			return nil, err
		}

		// If mutual like, create a match
		if mutualLike {
			var matchID int64
			err = tx.QueryRowContext(
				ctx,
				`INSERT INTO matches (user1_id, user2_id, created_at)
				VALUES ($1, $2, $3)
				RETURNING id`,
				userID, otherUserID, now,
			).Scan(&matchID)
			if err != nil {
				return nil, err
			}

			match = &models.Match{
				ID:        matchID,
				User1ID:   userID,
				User2ID:   otherUserID,
				CreatedAt: now,
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return match, nil
}

// GetMatches retrieves a user's matches
func (s *MatchingService) GetMatches(ctx context.Context, userID int64) ([]*models.Match, error) {
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT id, user1_id, user2_id, created_at
		FROM matches
		WHERE user1_id = $1 OR user2_id = $1
		ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []*models.Match
	for rows.Next() {
		var match models.Match
		if err := rows.Scan(&match.ID, &match.User1ID, &match.User2ID, &match.CreatedAt); err != nil {
			return nil, err
		}
		matches = append(matches, &match)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

// Helper function to get photos for a profile
func (s *MatchingService) getProfilePhotos(ctx context.Context, profileID int64) ([]models.Photo, error) {
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT id, profile_id, url, is_primary, created_at 
		FROM photos WHERE profile_id = $1 ORDER BY is_primary DESC, created_at ASC`,
		profileID,
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

	return photos, nil
}
