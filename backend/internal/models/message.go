package models

import "time"

// Message represents a message between matched users
type Message struct {
	ID        int64     `json:"id"`
	MatchID   int64     `json:"match_id"`
	SenderID  int64     `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageInput represents data needed to create a message
type MessageInput struct {
	Content string `json:"content" validate:"required"`
}
