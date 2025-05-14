package models

import "time"

// Match represents a match between two users
type Match struct {
	ID            int64     `json:"id"`
	User1ID       string    `json:"user1_id"`
	User2ID       string    `json:"user2_id"`
	CreatedAt     time.Time `json:"created_at"`
	LastMessageAt time.Time `json:"last_message_at"`
	User1LastRead time.Time `json:"user1_last_read"`
	User2LastRead time.Time `json:"user2_last_read"`
	OtherUser     *Profile  `json:"other_user,omitempty"` // Populated when retrieving matches
	LastMessage   *Message  `json:"last_message,omitempty"` // Last message in the conversation
	UnreadCount   int       `json:"unread_count,omitempty"` // Number of unread messages
}

// Swipe represents a user's swipe (like or skip) on another profile
type Swipe struct {
	ID        int64      `json:"id"`
	UserID    string     `json:"user_id"`
	ProfileID string     `json:"profile_id"`
	IsLike    bool       `json:"is_like"`
	Message   string     `json:"message,omitempty"` // Message attached to a like
	IsRose    bool       `json:"is_rose"`          // Whether this is a rose
	CreatedAt time.Time  `json:"created_at"`
	Profile   *Profile   `json:"profile,omitempty"` // Populated when retrieving likes
}

// SwipeInput represents the input for creating a swipe
type SwipeInput struct {
	ProfileID string `json:"profile_id"`
	IsLike    bool   `json:"is_like"`
	Message   string `json:"message,omitempty"`
	IsRose    bool   `json:"is_rose,omitempty"`
}

// Standout represents a standout profile recommendation
type Standout struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	ProfileID string    `json:"profile_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	IsActive  bool      `json:"is_active"`
	Profile   *Profile  `json:"profile,omitempty"` // Populated when retrieving standouts
}

// MatchWithProfile represents a match with the other user's profile
type MatchWithProfile struct {
	Match   Match   `json:"match"`
	Profile Profile `json:"profile"`
}

// MarkAsReadInput represents the input for marking messages as read
type MarkAsReadInput struct {
	MatchID int64 `json:"match_id"`
}
