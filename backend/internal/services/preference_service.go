package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// PreferenceService handles user preferences
type PreferenceService struct {
	BaseService
}

// NewPreferenceService creates a new preference service
func NewPreferenceService(db *sql.DB) *PreferenceService {
	return &PreferenceService{
		BaseService: BaseService{
			db: db,
		},
	}
}

// GetPreferences retrieves user preferences
func (s *PreferenceService) GetPreferences(ctx context.Context, userID string) (*models.Preference, error) {
	db := s.GetDB()

	var preference models.Preference
	var preferencesJSON []byte

	err := db.QueryRowContext(ctx, `
		SELECT id, user_id, preferred_gender, min_age, max_age, max_distance, preferences, created_at, updated_at
		FROM preferences 
		WHERE user_id = $1
	`, userID).Scan(
		&preference.ID, &preference.UserID, &preference.PreferredGender, &preference.MinAge, 
		&preference.MaxAge, &preference.MaxDistance, &preferencesJSON, &preference.CreatedAt, &preference.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Return default preferences
			return &models.Preference{
				UserID:          userID,
				PreferredGender: "",
				MinAge:          18,
				MaxAge:          100,
				MaxDistance:     50,
				Preferences:     make(map[string]interface{}),
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			}, nil
		}
		return nil, err
	}

	// Parse preferences JSON
	if preferencesJSON != nil {
		if err := json.Unmarshal(preferencesJSON, &preference.Preferences); err != nil {
			return nil, err
		}
	} else {
		preference.Preferences = make(map[string]interface{})
	}

	return &preference, nil
}

// CreateOrUpdatePreference creates or updates user preferences
func (s *PreferenceService) CreateOrUpdatePreference(ctx context.Context, userID string, preference *models.PreferenceInput) error {
	db := s.GetDB()

	// Validate input
	if preference.MinAge < 18 {
		return errors.New("minimum age must be at least 18")
	}
	if preference.MaxAge > 100 {
		return errors.New("maximum age must be at most 100")
	}
	if preference.MinAge > preference.MaxAge {
		return errors.New("minimum age cannot be greater than maximum age")
	}
	if preference.MaxDistance < 1 {
		return errors.New("maximum distance must be at least 1")
	}

	// Prepare preferences JSON
	preferencesJSON := []byte("{}")
	if preference.Preferences != nil {
		preferencesBytes, err := json.Marshal(preference.Preferences)
		if err != nil {
			return err
		}
		preferencesJSON = preferencesBytes
	}

	// Check if preferences already exist
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM preferences WHERE user_id = $1)", userID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// Update existing preferences
		_, err = db.ExecContext(ctx, `
			UPDATE preferences 
			SET preferred_gender = $1, min_age = $2, max_age = $3, max_distance = $4, 
			    preferences = $5, updated_at = NOW()
			WHERE user_id = $6
		`, preference.PreferredGender, preference.MinAge, preference.MaxAge, 
		   preference.MaxDistance, preferencesJSON, userID)
	} else {
		// Create new preferences
		_, err = db.ExecContext(ctx, `
			INSERT INTO preferences 
			(user_id, preferred_gender, min_age, max_age, max_distance, preferences)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, userID, preference.PreferredGender, preference.MinAge, preference.MaxAge, 
		   preference.MaxDistance, preferencesJSON)
	}

	return err
}

// UpdatePreferences updates user preferences
func (s *PreferenceService) UpdatePreferences(ctx context.Context, userID string, preferences map[string]interface{}) error {
	// TODO: Implement updating user preferences
	return nil
} 