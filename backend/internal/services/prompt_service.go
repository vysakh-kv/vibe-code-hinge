package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// PromptService handles prompt operations
type PromptService struct {
	BaseService
}

// NewPromptService creates a new PromptService
func NewPromptService(db *sql.DB) *PromptService {
	return &PromptService{
		BaseService: NewBaseService(db),
	}
}

// GetDefaultPrompts retrieves the list of default prompts
func (s *PromptService) GetDefaultPrompts(ctx context.Context) ([]models.Prompt, error) {
	rows, err := s.GetDB().QueryContext(
		ctx,
		`SELECT id, text 
		FROM prompts 
		ORDER BY id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prompts []models.Prompt
	for rows.Next() {
		var prompt models.Prompt
		if err := rows.Scan(&prompt.ID, &prompt.Text); err != nil {
			return nil, err
		}
		prompts = append(prompts, prompt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prompts, nil
}

// GetUserPrompts retrieves the prompts for a specific user
func (s *PromptService) GetUserPrompts(ctx context.Context, userID string) ([]map[string]interface{}, error) {
	// Get user's profile ID
	var profileID string
	err := s.GetDB().QueryRowContext(
		ctx,
		"SELECT id FROM profiles WHERE user_id = $1",
		userID,
	).Scan(&profileID)

	if err != nil {
		if err == sql.ErrNoRows {
			return []map[string]interface{}{}, nil
		}
		return nil, err
	}

	// Get profile prompts
	rows, err := s.GetDB().QueryContext(
		ctx,
		`SELECT pp.id, pp.prompt_id, p.text, pp.answer 
		FROM profile_prompts pp
		JOIN prompts p ON pp.prompt_id = p.id
		WHERE pp.profile_id = $1
		ORDER BY pp.created_at ASC`,
		profileID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prompts []map[string]interface{}
	for rows.Next() {
		var id, promptID int64
		var text, answer string
		if err := rows.Scan(&id, &promptID, &text, &answer); err != nil {
			return nil, err
		}
		prompts = append(prompts, map[string]interface{}{
			"id":        id,
			"prompt_id": promptID,
			"text":      text,
			"answer":    answer,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prompts, nil
}

// UpdateUserPrompt updates a user prompt
func (s *PromptService) UpdateUserPrompt(ctx context.Context, userID string, promptID string, response string) error {
	// Get user's profile ID
	var profileID string
	err := s.GetDB().QueryRowContext(
		ctx,
		"SELECT id FROM profiles WHERE user_id = $1",
		userID,
	).Scan(&profileID)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("profile not found")
		}
		return err
	}

	// Update or insert the prompt response
	_, err = s.GetDB().ExecContext(
		ctx,
		`INSERT INTO profile_prompts (profile_id, prompt_id, answer)
		VALUES ($1, $2, $3)
		ON CONFLICT (profile_id, prompt_id) 
		DO UPDATE SET answer = $3, updated_at = NOW()`,
		profileID, promptID, response,
	)
	
	return err
} 