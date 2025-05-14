package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// MessageService handles message operations
type MessageService struct {
	BaseService
	matchingService    *MatchingService
	notificationService *NotificationService
}

// NewMessageService creates a new message service
func NewMessageService(db *sql.DB) *MessageService {
	return &MessageService{
		BaseService: NewBaseService(db),
	}
}

// Initialize sets up references to other services to avoid circular dependencies
func (s *MessageService) Initialize(matchingService *MatchingService, notificationService *NotificationService) {
	s.matchingService = matchingService
	s.notificationService = notificationService
}

// SendMessage sends a message in a conversation
func (s *MessageService) SendMessage(ctx context.Context, userID string, input models.MessageInput) (*models.Message, error) {
	db := s.GetDB()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check if match exists and the user is part of it
	var user1ID, user2ID string
	var lastRead time.Time
	err = tx.QueryRowContext(ctx, `
		SELECT 
			m.user1_id, 
			m.user2_id,
			CASE 
				WHEN m.user1_id = $1 THEN m.user1_last_read 
				ELSE m.user2_last_read 
			END as last_read
		FROM matches m
		WHERE m.id = $2 AND (m.user1_id = $1 OR m.user2_id = $1)
	`, userID, input.MatchID).Scan(&user1ID, &user2ID, &lastRead)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("match not found or user not part of match")
		}
		return nil, err
	}

	// Get the recipient ID
	var recipientID string
	if user1ID == userID {
		recipientID = user2ID
	} else {
		recipientID = user1ID
	}

	// Create the message
	now := time.Now()
	var messageID int64
	err = tx.QueryRowContext(ctx, `
		INSERT INTO messages (match_id, sender_id, message, is_read, created_at)
		VALUES ($1, $2, $3, false, $4)
		RETURNING id
	`, input.MatchID, userID, input.Message, now).Scan(&messageID)
	if err != nil {
		return nil, err
	}

	// Update the match's last message time
	_, err = tx.ExecContext(ctx, `
		UPDATE matches
		SET last_message_at = $1
		WHERE id = $2
	`, now, input.MatchID)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// Create and return the message
	message := &models.Message{
		ID:        messageID,
		MatchID:   input.MatchID,
		SenderID:  userID,
		Message:   input.Message,
		IsRead:    false,
		CreatedAt: now,
	}

	// Send notification asynchronously
	if s.notificationService != nil {
		go func() {
			notifyCtx := context.Background()
			data := map[string]interface{}{
				"match_id":   input.MatchID,
				"sender_id":  userID,
				"message":    input.Message,
				"message_id": messageID,
			}
			_ = s.notificationService.SendNotification(notifyCtx, recipientID, "message", data)
		}()
	}

	return message, nil
}

// GetMessages retrieves messages for a conversation
func (s *MessageService) GetMessages(ctx context.Context, userID string, matchID int64, limit, offset int) ([]models.Message, error) {
	db := s.GetDB()

	// Verify match and user participation
	var exists bool
	err := db.QueryRowContext(ctx, `
		SELECT EXISTS(
			SELECT 1 FROM matches
			WHERE id = $1 AND (user1_id = $2 OR user2_id = $2)
		)
	`, matchID, userID).Scan(&exists)

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("match not found or user not part of match")
	}

	// Get messages with pagination
	rows, err := db.QueryContext(ctx, `
		SELECT id, match_id, sender_id, message, is_read, created_at
		FROM messages
		WHERE match_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, matchID, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the results
	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(
			&msg.ID,
			&msg.MatchID,
			&msg.SenderID,
			&msg.Message,
			&msg.IsRead,
			&msg.CreatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Mark the messages as read if they're from the other user
	// We do this in a separate transaction, and it's okay if it fails
	go func() {
		tx, err := db.Begin()
		if err != nil {
			return
		}
		defer tx.Rollback()

		// Update messages
		_, err = tx.Exec(`
			UPDATE messages
			SET is_read = true
			WHERE match_id = $1 AND sender_id != $2 AND is_read = false
		`, matchID, userID)
		if err != nil {
			return
		}

		// Update the user's last read time on the match
		_, err = tx.Exec(`
			UPDATE matches
			SET user1_last_read = NOW()
			WHERE id = $1 AND user1_id = $2
		`, matchID, userID)
		if err != nil {
			return
		}

		_, err = tx.Exec(`
			UPDATE matches
			SET user2_last_read = NOW()
			WHERE id = $1 AND user2_id = $2
		`, matchID, userID)
		if err != nil {
			return
		}

		_ = tx.Commit()
	}()

	return messages, nil
}

// MarkMessageAsRead marks a message as read
func (s *MessageService) MarkMessageAsRead(ctx context.Context, userID string, messageID int64) error {
	db := s.GetDB()

	// Check if the message exists and the user is the recipient
	var senderID string
	var matchID int64
	err := db.QueryRowContext(ctx, `
		SELECT m.sender_id, m.match_id
		FROM messages m
		JOIN matches mt ON m.match_id = mt.id
		WHERE m.id = $1 AND (mt.user1_id = $2 OR mt.user2_id = $2)
	`, messageID, userID).Scan(&senderID, &matchID)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("message not found or user not part of conversation")
		}
		return err
	}

	// Only mark as read if the user is not the sender
	if senderID == userID {
		return errors.New("cannot mark your own message as read")
	}

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Mark the message as read
	_, err = tx.ExecContext(ctx, `
		UPDATE messages
		SET is_read = true
		WHERE id = $1
	`, messageID)
	if err != nil {
		return err
	}

	// Update the user's last read time on the match
	if isUser1 := userID == senderID; isUser1 {
		_, err = tx.ExecContext(ctx, `
			UPDATE matches
			SET user1_last_read = NOW()
			WHERE id = $1
		`, matchID)
	} else {
		_, err = tx.ExecContext(ctx, `
			UPDATE matches
			SET user2_last_read = NOW()
			WHERE id = $1
		`, matchID)
	}
	if err != nil {
		return err
	}

	return tx.Commit()
}
