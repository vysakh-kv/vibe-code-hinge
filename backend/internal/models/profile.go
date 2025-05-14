package models

import "time"

// Profile represents a user's dating profile
type Profile struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Bio         string    `json:"bio"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Location    string    `json:"location"`
	Occupation  string    `json:"occupation"`
	Photos      []Photo   `json:"photos"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Photo represents a profile photo
type Photo struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profile_id"`
	URL       string    `json:"url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

// ProfileInput represents data needed to create or update a profile
type ProfileInput struct {
	Name        string   `json:"name" validate:"required"`
	Bio         string   `json:"bio"`
	DateOfBirth string   `json:"date_of_birth" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	Location    string   `json:"location"`
	Occupation  string   `json:"occupation"`
	Photos      []string `json:"photos"`
}
