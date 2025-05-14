package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// ProfileService handles profile-related business logic
type ProfileService struct {
	BaseService
}

// NewProfileService creates a new profile service
func NewProfileService(db *sql.DB) *ProfileService {
	return &ProfileService{
		BaseService: NewBaseService(db),
	}
}

// GetProfileByID retrieves a profile by its ID
func (s *ProfileService) GetProfileByID(ctx context.Context, id string) (*models.Profile, error) {
	db := s.GetDB()

	var profile models.Profile
	var userID string
	var vicesJSON, preferencesJSON []byte
	var dateOfBirth time.Time

	err := db.QueryRowContext(ctx, `
		SELECT p.id, p.user_id, p.name, p.bio, p.date_of_birth, 
		       p.gender, p.location, p.occupation, p.vices, p.preferences,
		       p.created_at, p.updated_at
		FROM profiles p
		WHERE p.id = $1
	`, id).Scan(
		&profile.ID, &userID, &profile.Name, &profile.Bio, &dateOfBirth,
		&profile.Gender, &profile.Location, &profile.Occupation, &vicesJSON, &preferencesJSON,
		&profile.CreatedAt, &profile.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	profile.UserID = userID
	profile.DateOfBirth = dateOfBirth.Format("2006-01-02")

	// Parse vices JSON
	if vicesJSON != nil {
		if err := json.Unmarshal(vicesJSON, &profile.Vices); err != nil {
			return nil, err
		}
	} else {
		profile.Vices = make(map[string]bool)
	}

	// Parse preferences JSON
	if preferencesJSON != nil {
		if err := json.Unmarshal(preferencesJSON, &profile.Preferences); err != nil {
			return nil, err
		}
	} else {
		profile.Preferences = make(map[string]interface{})
	}

	// Get profile photos
	photos, err := s.getProfilePhotos(ctx, id)
	if err != nil {
		return nil, err
	}
	profile.Photos = photos

	// Get profile prompts
	prompts, err := s.getProfilePrompts(ctx, id)
	if err != nil {
		return nil, err
	}
	profile.Prompts = prompts

	return &profile, nil
}

// CreateOrUpdateProfile creates or updates a profile
func (s *ProfileService) CreateOrUpdateProfile(ctx context.Context, userID string, profile *models.ProfileInput) (*models.Profile, error) {
	db := s.GetDB()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Check if profile already exists
	var profileID string
	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT id FROM profiles WHERE user_id = $1", userID).Scan(&profileID)
	if err == nil {
		exists = true
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	// Validate input
	if profile.Name == "" {
		return nil, errors.New("name is required")
	}
	if profile.DateOfBirth == "" {
		return nil, errors.New("date of birth is required")
	}
	if profile.Gender == "" {
		return nil, errors.New("gender is required")
	}

	// Parse date of birth
	dob, err := time.Parse("2006-01-02", profile.DateOfBirth)
	if err != nil {
		return nil, errors.New("invalid date of birth format, use YYYY-MM-DD")
	}

	// Calculate age
	age := time.Now().Year() - dob.Year()
	if time.Now().YearDay() < dob.YearDay() {
		age--
	}
	if age < 18 {
		return nil, errors.New("user must be at least 18 years old")
	}

	// Prepare preferences JSON
	preferencesJSON := []byte("{}")
	if profile.Preferences != nil {
		preferencesBytes, err := json.Marshal(profile.Preferences)
		if err != nil {
			return nil, err
		}
		preferencesJSON = preferencesBytes
	}

	// Prepare vices JSON
	vicesJSON := []byte("{}")
	if profile.Vices != nil {
		vicesBytes, err := json.Marshal(profile.Vices)
		if err != nil {
			return nil, err
		}
		vicesJSON = vicesBytes
	}

	// Create or update profile
	if exists {
		_, err = tx.ExecContext(ctx, `
			UPDATE profiles 
			SET name = $1, bio = $2, date_of_birth = $3, gender = $4, 
			    location = $5, occupation = $6, preferences = $7, updated_at = NOW(),
			    vices = $8
			WHERE id = $9
		`, profile.Name, profile.Bio, dob, profile.Gender,
			profile.Location, profile.Occupation, preferencesJSON, vicesJSON, profileID)
		if err != nil {
			return nil, err
		}
	} else {
		err = tx.QueryRowContext(ctx, `
			INSERT INTO profiles 
			(user_id, name, bio, date_of_birth, gender, location, occupation, preferences, vices)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id
		`, userID, profile.Name, profile.Bio, dob, profile.Gender,
			profile.Location, profile.Occupation, preferencesJSON, vicesJSON).Scan(&profileID)
		if err != nil {
			return nil, err
		}
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// Fetch the complete profile
	return s.GetProfileByID(ctx, profileID)
}

// GetProfilePrompts gets all prompts for a profile
func (s *ProfileService) GetProfilePrompts(ctx context.Context, profileID string) ([]models.ProfilePrompt, error) {
	return s.getProfilePrompts(ctx, profileID)
}

// Helper function to get profile prompts
func (s *ProfileService) getProfilePrompts(ctx context.Context, profileID string) ([]models.ProfilePrompt, error) {
	db := s.GetDB()

	rows, err := db.QueryContext(ctx, `
		SELECT pp.id, pp.profile_id, pp.prompt_id, p.text, pp.answer
		FROM profile_prompts pp
		JOIN prompts p ON pp.prompt_id = p.id
		WHERE pp.profile_id = $1
		ORDER BY pp.id
	`, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prompts []models.ProfilePrompt
	for rows.Next() {
		var prompt models.ProfilePrompt
		var promptID int64
		if err := rows.Scan(&prompt.ID, &prompt.ProfileID, &promptID, &prompt.Text, &prompt.Answer); err != nil {
			return nil, err
		}
		prompt.PromptID = promptID
		prompts = append(prompts, prompt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prompts, nil
}

// getProfilePhotos is an internal helper to get profile photos
func (s *ProfileService) getProfilePhotos(ctx context.Context, profileID string) ([]models.Photo, error) {
	rows, err := s.GetDB().QueryContext(
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

// SaveProfilePrompt saves a prompt response for a profile
func (s *ProfileService) SaveProfilePrompt(ctx context.Context, profileID string, input models.ProfilePromptInput) (*models.ProfilePrompt, error) {
	db := s.GetDB()

	// Check if prompt exists
	var promptText string
	err := db.QueryRowContext(ctx, "SELECT text FROM prompts WHERE id = $1", input.PromptID).Scan(&promptText)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("prompt not found")
		}
		return nil, err
	}

	// Check if profile exists
	var exists bool
	err = db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM profiles WHERE id = $1)", profileID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("profile not found")
	}

	// Check if profile prompt already exists
	var promptResponseID int64
	err = db.QueryRowContext(ctx, `
		SELECT id FROM profile_prompts 
		WHERE profile_id = $1 AND prompt_id = $2
	`, profileID, input.PromptID).Scan(&promptResponseID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		// Create new prompt response
		err = db.QueryRowContext(ctx, `
			INSERT INTO profile_prompts (profile_id, prompt_id, answer)
			VALUES ($1, $2, $3)
			RETURNING id
		`, profileID, input.PromptID, input.Answer).Scan(&promptResponseID)
		if err != nil {
			return nil, err
		}
	} else {
		// Update existing prompt response
		_, err = db.ExecContext(ctx, `
			UPDATE profile_prompts
			SET answer = $1, updated_at = NOW()
			WHERE id = $2
		`, input.Answer, promptResponseID)
		if err != nil {
			return nil, err
		}
	}

	// Return the saved prompt
	return &models.ProfilePrompt{
		ID:        promptResponseID,
		ProfileID: profileID,
		PromptID:  input.PromptID,
		Text:      promptText,
		Answer:    input.Answer,
	}, nil
}

// DeleteProfilePrompt deletes a prompt response for a profile
func (s *ProfileService) DeleteProfilePrompt(ctx context.Context, profileID string, promptID int64) error {
	db := s.GetDB()

	// Check if profile exists
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM profiles WHERE id = $1)", profileID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("profile not found")
	}

	// Delete the prompt response
	result, err := db.ExecContext(ctx, `
		DELETE FROM profile_prompts
		WHERE profile_id = $1 AND prompt_id = $2
	`, profileID, promptID)
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("prompt response not found")
	}

	return nil
}

// CompleteOnboarding marks the profile as having completed onboarding
func (s *ProfileService) CompleteOnboarding(ctx context.Context, profileID string) error {
	db := s.GetDB()

	// Check if profile exists
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM profiles WHERE id = $1)", profileID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("profile not found")
	}

	// Mark profile as having completed onboarding
	_, err = db.ExecContext(ctx, `
		UPDATE profiles
		SET onboarding_completed = true, updated_at = NOW()
		WHERE id = $1
	`, profileID)
	if err != nil {
		return err
	}

	return nil
}

// GetProfileByUserID retrieves a profile by user ID
func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID string) (*models.Profile, error) {
	var profileID string
	err := s.GetDB().QueryRowContext(
		ctx,
		"SELECT id FROM profiles WHERE user_id = $1",
		userID,
	).Scan(&profileID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	return s.GetProfileByID(ctx, profileID)
}
