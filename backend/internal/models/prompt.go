package models

import "time"

// PromptResponse represents a user's response to a prompt
type PromptResponse struct {
	ID        int64     `json:"id"`
	ProfileID string    `json:"profile_id"`
	PromptID  int64     `json:"prompt_id"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} 