package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// NotificationService handles notifications and SSE operations
type NotificationService struct {
	BaseService
	clients    map[string]map[chan []byte]bool
	clientsMux sync.RWMutex
}

// NewNotificationService creates a new NotificationService
func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{
		BaseService: NewBaseService(db),
		clients:     make(map[string]map[chan []byte]bool),
		clientsMux:  sync.RWMutex{},
	}
}

// StreamEvents streams events to the client using SSE
func (s *NotificationService) StreamEvents(w http.ResponseWriter, r *http.Request, userID string) error {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for this client
	messageChan := make(chan []byte)

	// Register client
	s.clientsMux.Lock()
	if _, exists := s.clients[userID]; !exists {
		s.clients[userID] = make(map[chan []byte]bool)
	}
	s.clients[userID][messageChan] = true
	s.clientsMux.Unlock()

	// Ensure client is removed when connection is closed
	defer func() {
		s.clientsMux.Lock()
		delete(s.clients[userID], messageChan)
		if len(s.clients[userID]) == 0 {
			delete(s.clients, userID)
		}
		s.clientsMux.Unlock()
		close(messageChan)
	}()

	// Create channel to notify when client disconnects
	notify := r.Context().Done()
	go func() {
		<-notify
		// Connection is closed, clean up
		s.clientsMux.Lock()
		delete(s.clients[userID], messageChan)
		if len(s.clients[userID]) == 0 {
			delete(s.clients, userID)
		}
		s.clientsMux.Unlock()
		close(messageChan)
	}()

	// Send test event
	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	// Send "connected" event
	initialEvent := models.NotificationEvent{
		Type:      "connected",
		UserID:    userID,
		Message:   "Connected to event stream",
		Timestamp: time.Now(),
	}
	initialEventData, _ := json.Marshal(initialEvent)
	fmt.Fprintf(w, "data: %s\n\n", initialEventData)
	w.(http.Flusher).Flush()

	// Keep connection alive and send messages when they arrive
	for {
		select {
		case <-pingTicker.C:
			// Send ping to keep connection alive
			fmt.Fprintf(w, "event: ping\ndata: {\"time\": \"%s\"}\n\n", time.Now().Format(time.RFC3339))
			w.(http.Flusher).Flush()
		case msg, ok := <-messageChan:
			if !ok {
				// Channel closed, return
				return nil
			}
			// Send message to client
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-notify:
			// Client disconnected
			return nil
		}
	}
}

// SendNotification sends a notification to a user
func (s *NotificationService) SendNotification(ctx context.Context, userID string, notificationType string, data map[string]interface{}) error {
	// Store notification in database
	now := time.Now()
	var targetID int64
	if val, ok := data["target_id"]; ok {
		if tID, ok := val.(int64); ok {
			targetID = tID
		}
	}
	
	message, _ := data["message"].(string)

	// Create notification record
	_, err := s.GetDB().ExecContext(
		ctx,
		`INSERT INTO notifications (user_id, type, target_id, message, is_read, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		userID, notificationType, targetID, message, false, now,
	)
	if err != nil {
		log.Printf("Failed to save notification: %v", err)
		// Continue even if DB save fails - we can still deliver the in-memory notification
	}

	// Create event for SSE
	event := models.NotificationEvent{
		Type:      notificationType,
		UserID:    userID,
		TargetID:  targetID,
		Message:   message,
		Timestamp: now,
	}

	// When type is "message", create a MessageEvent instead
	var eventData []byte
	var err2 error
	
	if notificationType == "message" {
		// For message events, get the data we need
		var matchID int64
		var senderID string
		var messageText string
		var messageID int64
		
		if val, ok := data["match_id"].(int64); ok {
			matchID = val
		}
		if val, ok := data["sender_id"].(string); ok {
			senderID = val
		}
		if val, ok := data["message"].(string); ok {
			messageText = val
		}
		if val, ok := data["message_id"].(int64); ok {
			messageID = val
		}
		
		messageEvent := models.MessageEvent{
			Type:      "message",
			MessageID: messageID,
			MatchID:   matchID,
			SenderID:  senderID,
			Message:   messageText,
			CreatedAt: now,
		}
		eventData, err2 = json.Marshal(messageEvent)
	} else {
		eventData, err2 = json.Marshal(event)
	}
	
	if err2 != nil {
		return err2
	}

	// Send to connected clients
	s.clientsMux.RLock()
	if clients, exists := s.clients[userID]; exists {
		for clientChan := range clients {
			select {
			case clientChan <- eventData:
				// Message sent
			default:
				// Channel buffer full, skip this message
				log.Printf("Channel buffer full for user %s, skipping message", userID)
			}
		}
	}
	s.clientsMux.RUnlock()

	return nil
}

// GetUnreadNotificationCount gets the count of unread notifications for a user
func (s *NotificationService) GetUnreadNotificationCount(ctx context.Context, userID string) (int, error) {
	var count int
	err := s.GetDB().QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false",
		userID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetNotifications gets notifications for a user
func (s *NotificationService) GetNotifications(ctx context.Context, userID string, limit, offset int) ([]map[string]interface{}, error) {
	rows, err := s.GetDB().QueryContext(
		ctx,
		`SELECT id, type, target_id, message, is_read, created_at
		FROM notifications
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []map[string]interface{}
	for rows.Next() {
		var id, targetID int64
		var notificationType, message string
		var isRead bool
		var createdAt time.Time

		err := rows.Scan(&id, &notificationType, &targetID, &message, &isRead, &createdAt)
		if err != nil {
			return nil, err
		}

		notification := map[string]interface{}{
			"id":         id,
			"type":       notificationType,
			"target_id":  targetID,
			"message":    message,
			"is_read":    isRead,
			"created_at": createdAt,
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

// MarkNotificationAsRead marks a notification as read
func (s *NotificationService) MarkNotificationAsRead(ctx context.Context, userID string, notificationID int64) error {
	result, err := s.GetDB().ExecContext(
		ctx,
		"UPDATE notifications SET is_read = true WHERE id = $1 AND user_id = $2",
		notificationID, userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("notification not found or not owned by user")
	}

	return nil
}

// MarkAllNotificationsAsRead marks all notifications as read for a user
func (s *NotificationService) MarkAllNotificationsAsRead(ctx context.Context, userID string) error {
	_, err := s.GetDB().ExecContext(
		ctx,
		"UPDATE notifications SET is_read = true WHERE user_id = $1 AND is_read = false",
		userID,
	)
	return err
} 