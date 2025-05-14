package services

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"
)

// FeedService handles feed and discovery operations
type FeedService struct {
	BaseService
	profileService *ProfileService
}

// NewFeedService creates a new FeedService
func NewFeedService(db *sql.DB) *FeedService {
	return &FeedService{
		BaseService:    NewBaseService(db),
		profileService: NewProfileService(db),
	}
}

// GetFeed retrieves profiles for the main feed
func (s *FeedService) GetFeed(ctx context.Context, userID string, limit int, offset int) ([]map[string]interface{}, error) {
	// Get user preferences
	preferenceService := NewPreferenceService(s.GetDB())
	preferences, err := preferenceService.GetPreferences(ctx, userID)
	if err != nil {
		log.Printf("Failed to get preferences, using defaults: %v", err)
		// Default preferences will be applied
	}

	// Build query based on preferences
	query := `
		SELECT p.id
		FROM profiles p
		LEFT JOIN swipes s ON p.id = s.profile_id AND s.user_id = $1
		WHERE s.id IS NULL AND p.user_id != $1
	`
	args := []interface{}{userID}
	argCount := 1

	// Apply gender filter if specified
	if preferences != nil && preferences.PreferredGender != "all" {
		argCount++
		query += ` AND p.gender = $` + strconv.Itoa(argCount)
		args = append(args, preferences.PreferredGender)
	}

	// Apply age filter if specified
	if preferences != nil && preferences.MinAge > 0 {
		argCount++
		query += ` AND EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.date_of_birth)) >= $` + strconv.Itoa(argCount)
		args = append(args, preferences.MinAge)
	}

	if preferences != nil && preferences.MaxAge > 0 {
		argCount++
		query += ` AND EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.date_of_birth)) <= $` + strconv.Itoa(argCount)
		args = append(args, preferences.MaxAge)
	}

	// Limit and offset
	query += ` ORDER BY RANDOM() LIMIT $` + strconv.Itoa(argCount+1) + ` OFFSET $` + strconv.Itoa(argCount+2)
	args = append(args, limit, offset)

	// Execute query
	rows, err := s.GetDB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect profile IDs
	var profileIDs []string
	for rows.Next() {
		var profileID string
		if err := rows.Scan(&profileID); err != nil {
			return nil, err
		}
		profileIDs = append(profileIDs, profileID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Get full profiles
	var profiles []map[string]interface{}
	for _, profileID := range profileIDs {
		profile, err := s.profileService.GetProfileByID(ctx, profileID)
		if err != nil {
			log.Printf("Failed to get profile %s: %v", profileID, err)
			continue
		}

		// Convert to map
		profileMap := map[string]interface{}{
			"id":          profile.ID,
			"name":        profile.Name,
			"gender":      profile.Gender,
			"bio":         profile.Bio,
			"age":         profile.Age(),
			"location":    profile.Location,
			"occupation":  profile.Occupation,
			"photos":      profile.Photos,
			"prompts":     profile.Prompts,
		}

		profiles = append(profiles, profileMap)
	}

	return profiles, nil
}

// GetStandouts retrieves standout profiles
func (s *FeedService) GetStandouts(ctx context.Context, userID string, limit int) ([]map[string]interface{}, error) {
	// First, clean up expired standouts
	_, err := s.GetDB().ExecContext(
		ctx,
		`UPDATE standouts SET is_active = false WHERE user_id = $1 AND expires_at < NOW()`,
		userID,
	)
	if err != nil {
		log.Printf("Failed to clean up expired standouts: %v", err)
	}

	// Check if we have enough active standouts
	var activeCount int
	err = s.GetDB().QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM standouts WHERE user_id = $1 AND is_active = true`,
		userID,
	).Scan(&activeCount)
	if err != nil {
		return nil, err
	}

	// If we don't have enough active standouts, generate new ones
	if activeCount < limit {
		neededCount := limit - activeCount
		err = s.generateStandouts(ctx, userID, neededCount)
		if err != nil {
			log.Printf("Failed to generate standouts: %v", err)
		}
	}

	// Retrieve active standouts
	rows, err := s.GetDB().QueryContext(
		ctx,
		`SELECT s.profile_id 
		FROM standouts s
		WHERE s.user_id = $1 AND s.is_active = true
		ORDER BY s.created_at DESC
		LIMIT $2`,
		userID, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect profile IDs
	var profileIDs []string
	for rows.Next() {
		var profileID string
		if err := rows.Scan(&profileID); err != nil {
			return nil, err
		}
		profileIDs = append(profileIDs, profileID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Get full profiles
	var standouts []map[string]interface{}
	for _, profileID := range profileIDs {
		profile, err := s.profileService.GetProfileByID(ctx, profileID)
		if err != nil {
			log.Printf("Failed to get profile %s: %v", profileID, err)
			continue
		}

		// Convert to map
		standoutMap := map[string]interface{}{
			"id":          profile.ID,
			"name":        profile.Name,
			"gender":      profile.Gender,
			"bio":         profile.Bio,
			"age":         profile.Age(),
			"location":    profile.Location,
			"occupation":  profile.Occupation,
			"photos":      profile.Photos,
			"prompts":     profile.Prompts,
			"standout_reason": "Popular profile", // Placeholder, would be determined by algorithm
		}

		standouts = append(standouts, standoutMap)
	}

	return standouts, nil
}

// generateStandouts creates new standout recommendations
func (s *FeedService) generateStandouts(ctx context.Context, userID string, count int) error {
	// Get user preferences
	preferenceService := NewPreferenceService(s.GetDB())
	preferences, err := preferenceService.GetPreferences(ctx, userID)
	if err != nil {
		log.Printf("Failed to get preferences, using defaults: %v", err)
		// Default preferences will be applied
	}

	// Find profiles not already in standouts or swiped
	query := `
		SELECT p.id
		FROM profiles p
		LEFT JOIN standouts s ON p.id = s.profile_id AND s.user_id = $1
		LEFT JOIN swipes sw ON p.id = sw.profile_id AND sw.user_id = $1
		WHERE s.id IS NULL AND sw.id IS NULL AND p.user_id != $1
	`
	args := []interface{}{userID}
	argCount := 1

	// Apply gender filter if specified
	if preferences != nil && preferences.PreferredGender != "all" {
		argCount++
		query += ` AND p.gender = $` + strconv.Itoa(argCount)
		args = append(args, preferences.PreferredGender)
	}

	// Apply age filter if specified
	if preferences != nil && preferences.MinAge > 0 {
		argCount++
		query += ` AND EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.date_of_birth)) >= $` + strconv.Itoa(argCount)
		args = append(args, preferences.MinAge)
	}

	if preferences != nil && preferences.MaxAge > 0 {
		argCount++
		query += ` AND EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.date_of_birth)) <= $` + strconv.Itoa(argCount)
		args = append(args, preferences.MaxAge)
	}

	// Order by random and limit
	query += ` ORDER BY RANDOM() LIMIT $` + strconv.Itoa(argCount+1)
	args = append(args, count)

	// Execute query
	rows, err := s.GetDB().QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Create standouts
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	tx, err := s.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for rows.Next() {
		var profileID string
		if err := rows.Scan(&profileID); err != nil {
			return err
		}

		_, err := tx.ExecContext(
			ctx,
			`INSERT INTO standouts (user_id, profile_id, created_at, expires_at, is_active)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (user_id, profile_id) DO UPDATE SET is_active = $5, expires_at = $4`,
			userID, profileID, now, expiresAt, true,
		)
		if err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return tx.Commit()
} 