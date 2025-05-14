package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// MessageService handles message-related business logic
type MessageService struct {
	db *sql.DB
}

// NewMessageService creates a new message service
func NewMessageService(db *sql.DB) *MessageService {
	return &MessageService{
		db: db,
	}
}

// GetMessages retrieves messages for a match
func (s *MessageService) GetMessages(ctx context.Context, matchID int64, userID int64) ([]*models.Message, error) {
	// Verify user is part of the match
	var isUserPartOfMatch bool
	err := s.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM matches WHERE id = $1 AND (user1_id = $2 OR user2_id = $2))",
		matchID, userID,
	).Scan(&isUserPartOfMatch)

	if err != nil {
		return nil, err
	}

	if !isUserPartOfMatch {
		return nil, errors.New("user is not part of this match")
	}

	// Get messages
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT id, match_id, sender_id, content, created_at
		FROM messages
		WHERE match_id = $1
		ORDER BY created_at ASC`,
		matchID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.MatchID, &message.SenderID, &message.Content, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

// CreateMessage creates a new message
func (s *MessageService) CreateMessage(ctx context.Context, matchID int64, senderID int64, input models.MessageInput) (*models.Message, error) {
	// Verify user is part of the match
	var isUserPartOfMatch bool
	err := s.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM matches WHERE id = $1 AND (user1_id = $2 OR user2_id = $2))",
		matchID, senderID,
	).Scan(&isUserPartOfMatch)

	if err != nil {
		return nil, err
	}

	if !isUserPartOfMatch {
		return nil, errors.New("user is not part of this match")
	}

	// Create message
	var message models.Message
	now := time.Now()

	err = s.db.QueryRowContext(
		ctx,
		`INSERT INTO messages (match_id, sender_id, content, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, match_id, sender_id, content, created_at`,
		matchID, senderID, input.Content, now,
	).Scan(&message.ID, &message.MatchID, &message.SenderID, &message.Content, &message.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &message, nil
}
