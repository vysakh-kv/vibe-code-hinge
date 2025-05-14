package models

import "time"

// Match represents a match between two users
type Match struct {
	ID        int64     `json:"id"`
	User1ID   int64     `json:"user1_id"`
	User2ID   int64     `json:"user2_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Swipe represents a user's swipe on another profile
type Swipe struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	ProfileID int64     `json:"profile_id"`
	IsLike    bool      `json:"is_like"`
	CreatedAt time.Time `json:"created_at"`
}

// SwipeInput represents the data needed to create a swipe
type SwipeInput struct {
	ProfileID int64 `json:"profile_id" validate:"required"`
	IsLike    bool  `json:"is_like" validate:"required"`
}
