package models

import (
	"time"
)

// Profile represents a user profile
type Profile struct {
	ID         string            `json:"id"`
	UserID     string            `json:"user_id"`
	Name       string            `json:"name"`
	Bio        string            `json:"bio,omitempty"`
	DateOfBirth string            `json:"date_of_birth"`
	Gender     string            `json:"gender"`
	Location   string            `json:"location,omitempty"`
	Occupation string            `json:"occupation,omitempty"`
	Vices      map[string]bool   `json:"vices,omitempty"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
	Photos     []Photo           `json:"photos,omitempty"`
	Prompts    []ProfilePrompt   `json:"prompts,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	OnboardingCompleted bool     `json:"onboarding_completed"`
}

// Photo represents a profile photo
type Photo struct {
	ID        int64     `json:"id"`
	ProfileID string    `json:"profile_id"`
	URL       string    `json:"url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

// ProfileInput represents input for creating or updating a profile
type ProfileInput struct {
	Name       string            `json:"name"`
	Bio        string            `json:"bio,omitempty"`
	DateOfBirth string            `json:"date_of_birth"`
	Gender     string            `json:"gender"`
	Location   string            `json:"location,omitempty"`
	Occupation string            `json:"occupation,omitempty"`
	Photos     []string          `json:"photos,omitempty"`
	Vices      map[string]bool   `json:"vices,omitempty"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
}

// ProfilePrompt represents a profile prompt with its answer
type ProfilePrompt struct {
	ID        int64     `json:"id"`
	ProfileID string    `json:"profile_id"`
	PromptID  int64     `json:"prompt_id"`
	Text      string    `json:"text"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Prompt represents a prompt template
type Prompt struct {
	ID    int64  `json:"id"`
	Text  string `json:"text"`
}

// ProfilePromptInput represents input for a profile prompt
type ProfilePromptInput struct {
	PromptID int64  `json:"prompt_id"`
	Answer   string `json:"answer"`
}

// OnboardingInput represents complete onboarding data
type OnboardingInput struct {
	Profile     ProfileInput       `json:"profile"`
	Preferences map[string]interface{} `json:"preferences"`
	Prompts     []ProfilePromptInput   `json:"prompts"`
}

// Age calculates the age of a profile based on date of birth
func (p *Profile) Age() int {
	today := time.Now()
	
	// Parse the date of birth string
	dob, err := time.Parse("2006-01-02", p.DateOfBirth)
	if err != nil {
		// Return 0 if date parsing fails
		return 0
	}
	
	// Calculate basic age
	age := today.Year() - dob.Year()
	
	// Adjust age if birthday hasn't occurred yet this year
	birthdayThisYear := time.Date(today.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)
	if birthdayThisYear.After(today) {
		age--
	}
	
	return age
}
