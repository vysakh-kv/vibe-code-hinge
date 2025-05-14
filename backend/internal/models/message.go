package models

import "time"

// Message represents a message in a conversation
type Message struct {
	ID        int64     `json:"id"`
	MatchID   int64     `json:"match_id"`
	SenderID  string    `json:"sender_id"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	Sender    *Profile  `json:"sender,omitempty"` // Populated when retrieving messages
}

// MessageInput represents the input for creating a message
type MessageInput struct {
	MatchID  int64  `json:"match_id"`
	Message  string `json:"message"`
}

// ConversationResponse represents a conversation with messages
type ConversationResponse struct {
	Match    Match     `json:"match"`
	Messages []Message `json:"messages"`
}

// MessageEvent represents a message event for SSE
type MessageEvent struct {
	Type      string    `json:"type"`
	MessageID int64     `json:"message_id"`
	MatchID   int64     `json:"match_id"`
	SenderID  string    `json:"sender_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// NotificationEvent represents a notification event for SSE
type NotificationEvent struct {
	Type        string    `json:"type"`
	UserID      string    `json:"user_id"`
	TargetID    int64     `json:"target_id,omitempty"`
	Message     string    `json:"message,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}
