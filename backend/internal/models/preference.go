package models

import (
	"time"
)

// Preference represents user preferences for matching
type Preference struct {
	ID              int64                  `json:"id"`
	UserID          string                 `json:"user_id"`
	PreferredGender string                 `json:"preferred_gender"`
	MinAge          int                    `json:"min_age"`
	MaxAge          int                    `json:"max_age"`
	MaxDistance     int                    `json:"max_distance"`
	Preferences     map[string]interface{} `json:"preferences"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// PreferenceInput represents input for creating or updating preferences
type PreferenceInput struct {
	PreferredGender string                 `json:"preferred_gender"`
	MinAge          int                    `json:"min_age"`
	MaxAge          int                    `json:"max_age"`
	MaxDistance     int                    `json:"max_distance"`
	Preferences     map[string]interface{} `json:"preferences"`
} 