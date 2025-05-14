package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// MatchingService handles matching operations
type MatchingService struct {
	BaseService
	profileService      *ProfileService
	notificationService *NotificationService
}

// NewMatchingService creates a new matching service
func NewMatchingService(db *sql.DB) *MatchingService {
	return &MatchingService{
		BaseService:     NewBaseService(db),
		profileService:  NewProfileService(db),
	}
}

// SetNotificationService sets the notification service
func (s *MatchingService) SetNotificationService(notificationService *NotificationService) {
	s.notificationService = notificationService
}

// GetDiscoverProfiles retrieves profiles for the discover feed
func (s *MatchingService) GetDiscoverProfiles(ctx context.Context, userID string, limit int) ([]map[string]interface{}, error) {
	// Create and use FeedService for this
	feedService := NewFeedService(s.GetDB())
	return feedService.GetFeed(ctx, userID, limit, 0)
}

// CreateSwipe creates a swipe and checks for a match
func (s *MatchingService) CreateSwipe(ctx context.Context, userID string, profileID string, isLike bool) (*models.Match, error) {
	db := s.GetDB()

	// Get profile's user_id
	var profileUserID string
	err := db.QueryRowContext(ctx, `SELECT user_id FROM profiles WHERE id = $1`, profileID).Scan(&profileUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Record the swipe
	_, err = tx.ExecContext(ctx, `
		INSERT INTO swipes (user_id, profile_id, is_like, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, profile_id) DO UPDATE SET is_like = $3
	`, userID, profileID, isLike, time.Now())
	
	if err != nil {
		return nil, err
	}

	// If it's not a like, no need to check for a match
	if !isLike {
		err = tx.Commit()
		return nil, err
	}

	// Check if the other user has already liked this user (mutual like = match)
	var mutualLike bool
	err = tx.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM swipes 
			WHERE user_id = $1 AND profile_id = $2 AND is_like = true
		)
	`, profileUserID, userID).Scan(&mutualLike)

	if err != nil {
		return nil, err
	}

	// If there's no mutual like, just return
	if !mutualLike {
		err = tx.Commit()
		return nil, err
	}

	// Create a match
	var matchID int64
	now := time.Now()

	// Ensure consistent ordering of user IDs
	var user1ID, user2ID string
	if userID < profileUserID {
		user1ID, user2ID = userID, profileUserID
	} else {
		user1ID, user2ID = profileUserID, userID
	}

	// Check if match already exists
	var matchExists bool
	err = tx.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM matches 
			WHERE (user1_id = $1 AND user2_id = $2) OR (user1_id = $2 AND user2_id = $1)
		)
	`, user1ID, user2ID).Scan(&matchExists)

	if err != nil {
		return nil, err
	}

	if !matchExists {
		// Create new match
		err = tx.QueryRowContext(ctx, `
			INSERT INTO matches (user1_id, user2_id, created_at, last_message_at, user1_last_read, user2_last_read)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
		`, user1ID, user2ID, now, now, now, now).Scan(&matchID)
		
		if err != nil {
			return nil, err
		}
	} else {
		// Get existing match ID
		err = tx.QueryRowContext(ctx, `
			SELECT id FROM matches 
			WHERE (user1_id = $1 AND user2_id = $2) OR (user1_id = $2 AND user2_id = $1)
		`, user1ID, user2ID).Scan(&matchID)
		
		if err != nil {
			return nil, err
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// Send notification if needed
	if s.notificationService != nil && !matchExists {
		// Get profiles for match notification
		profile1, _ := s.profileService.GetProfileByUserID(ctx, user1ID)
		profile2, _ := s.profileService.GetProfileByUserID(ctx, user2ID)

		// Send notification to both users
		if profile1 != nil && profile2 != nil {
			go func() {
				matchCtx := context.Background()
				s.notificationService.SendNotification(matchCtx, user1ID, "match", map[string]interface{}{
					"match_id":  matchID,
					"profile_id": profile2.ID,
					"message":   "You matched with " + profile2.Name,
				})
				s.notificationService.SendNotification(matchCtx, user2ID, "match", map[string]interface{}{
					"match_id":  matchID,
					"profile_id": profile1.ID,
					"message":   "You matched with " + profile1.Name,
				})
			}()
		}
	}

	// Return the match details
	return &models.Match{
		ID:            matchID,
		User1ID:       user1ID,
		User2ID:       user2ID,
		CreatedAt:     now,
		LastMessageAt: now,
	}, nil
}

// GetMatches retrieves all matches for a user
func (s *MatchingService) GetMatches(ctx context.Context, userID string) ([]models.MatchWithProfile, error) {
	db := s.GetDB()

	rows, err := db.QueryContext(ctx, `
		SELECT m.id, 
			m.user1_id, 
			m.user2_id, 
			m.created_at, 
			m.last_message_at,
			CASE 
				WHEN m.user1_id = $1 THEN m.user1_last_read 
				ELSE m.user2_last_read 
			END AS last_read
		FROM matches m
		WHERE m.user1_id = $1 OR m.user2_id = $1
		ORDER BY m.last_message_at DESC
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.MatchWithProfile
	for rows.Next() {
		var match models.Match
		var lastRead time.Time

		if err := rows.Scan(
			&match.ID,
			&match.User1ID, 
			&match.User2ID, 
			&match.CreatedAt, 
			&match.LastMessageAt,
			&lastRead,
		); err != nil {
			return nil, err
		}

		// Determine partner's ID
		var partnerID string
		if match.User1ID == userID {
			partnerID = match.User2ID
		} else {
			partnerID = match.User1ID
		}

		// Get the partner's profile
		profile, err := s.profileService.GetProfileByUserID(ctx, partnerID)
		if err != nil {
			continue
		}

		// Get unread count
		var unreadCount int
		err = db.QueryRowContext(ctx, `
			SELECT COUNT(*) 
			FROM messages 
			WHERE match_id = $1 AND sender_id = $2 AND created_at > $3
		`, match.ID, partnerID, lastRead).Scan(&unreadCount)
		
		if err != nil {
			unreadCount = 0
		}

		// Get last message if any
		var lastMessage models.Message
		err = db.QueryRowContext(ctx, `
			SELECT id, match_id, sender_id, message, is_read, created_at 
			FROM messages 
			WHERE match_id = $1 
			ORDER BY created_at DESC LIMIT 1
		`, match.ID).Scan(
			&lastMessage.ID,
			&lastMessage.MatchID,
			&lastMessage.SenderID,
			&lastMessage.Message,
			&lastMessage.IsRead,
			&lastMessage.CreatedAt,
		)

		if err != nil && err != sql.ErrNoRows {
			continue
		}

		if err == nil {
			match.LastMessage = &lastMessage
		}
		
		match.UnreadCount = unreadCount

		// Create MatchWithProfile
		matchWithProfile := models.MatchWithProfile{
			Match:   match,
			Profile: *profile,
		}

		matches = append(matches, matchWithProfile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

// MarkAsRead marks a match as read by the user
func (s *MatchingService) MarkAsRead(ctx context.Context, userID string, input models.MarkAsReadInput) error {
	db := s.GetDB()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// First check if the match exists and involves this user
	var user1ID, user2ID string
	err = tx.QueryRowContext(ctx, `
		SELECT user1_id, user2_id 
		FROM matches 
		WHERE id = $1
	`, input.MatchID).Scan(&user1ID, &user2ID)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("match not found")
		}
		return err
	}

	// Check if user is part of this match
	if userID != user1ID && userID != user2ID {
		return errors.New("user not part of this match")
	}

	// Update the last_read timestamp
	now := time.Now()
	if userID == user1ID {
		_, err = tx.ExecContext(ctx, `
			UPDATE matches 
			SET user1_last_read = $1 
			WHERE id = $2
		`, now, input.MatchID)
	} else {
		_, err = tx.ExecContext(ctx, `
			UPDATE matches 
			SET user2_last_read = $1 
			WHERE id = $2
		`, now, input.MatchID)
	}

	if err != nil {
		return err
	}

	// Also mark all messages as read
	_, err = tx.ExecContext(ctx, `
		UPDATE messages 
		SET is_read = true 
		WHERE match_id = $1 AND sender_id != $2
	`, input.MatchID, userID)

	if err != nil {
		return err
	}

	// Commit transaction
	return tx.Commit()
}

// GetConversation retrieves a conversation (match + messages)
func (s *MatchingService) GetConversation(ctx context.Context, userID string, matchID int64) (*models.ConversationResponse, error) {
	// Get the match
	match, err := s.GetMatch(ctx, userID, matchID)
	if err != nil {
		return nil, err
	}

	// Mark as read if there are unread messages
	if match.UnreadCount > 0 {
		err = s.MarkAsRead(ctx, userID, models.MarkAsReadInput{
			MatchID: matchID,
		})
		if err != nil {
			// Don't fail the request if marking as read fails
			// Just log it and continue
			// log.Printf("Failed to mark messages as read: %v", err)
		}
	}

	// Get the messages
	rows, err := s.GetDB().QueryContext(
		ctx,
		`SELECT 
			m.id, 
			m.match_id, 
			m.sender_id, 
			m.message, 
			m.is_read,
			m.created_at
		FROM messages m
		WHERE m.match_id = $1
		ORDER BY m.created_at ASC`,
		matchID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message

		err := rows.Scan(
			&msg.ID,
			&msg.MatchID,
			&msg.SenderID,
			&msg.Message,
			&msg.IsRead,
			&msg.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Get the partner's user ID
	var partnerID string
	if match.User1ID == userID {
		partnerID = match.User2ID
	} else {
		partnerID = match.User1ID
	}

	// Get the partner's profile
	profile, err := s.profileService.GetProfileByUserID(ctx, partnerID)
	if err != nil {
		return nil, err
	}

	// Add partner profile to match
	match.OtherUser = profile

	// Return the conversation
	return &models.ConversationResponse{
		Match:    *match,
		Messages: messages,
	}, nil
}

// GetMatch retrieves a specific match by ID
func (s *MatchingService) GetMatch(ctx context.Context, userID string, matchID int64) (*models.Match, error) {
	// Get the match
	var match models.Match
	var lastMessageAt sql.NullTime
	var user1LastRead, user2LastRead sql.NullTime

	err := s.GetDB().QueryRowContext(
		ctx,
		`SELECT
			m.id,
			m.user1_id,
			m.user2_id,
			m.created_at,
			m.last_message_at,
			m.user1_last_read,
			m.user2_last_read,
			COALESCE(m.unread_count, 0) as unread_count
		FROM matches m
		WHERE m.id = $1 AND (m.user1_id = $2 OR m.user2_id = $2)`,
		matchID, userID,
	).Scan(
		&match.ID,
		&match.User1ID,
		&match.User2ID,
		&match.CreatedAt,
		&lastMessageAt,
		&user1LastRead,
		&user2LastRead,
		&match.UnreadCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	if lastMessageAt.Valid {
		match.LastMessageAt = lastMessageAt.Time
	}

	if user1LastRead.Valid {
		match.User1LastRead = user1LastRead.Time
	}

	if user2LastRead.Valid {
		match.User2LastRead = user2LastRead.Time
	}

	return &match, nil
} 