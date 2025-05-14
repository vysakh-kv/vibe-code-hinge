package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vibe-code-hinge/backend/internal/services"
)

// NotificationHandler handles notification-related routes and SSE
type NotificationHandler struct {
	notificationService *services.NotificationService
}

// NewNotificationHandler creates a new notification handler
func NewNotificationHandler(notificationService *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// MessageEvents handles SSE for message events
func (h *NotificationHandler) MessageEvents(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create notification channel
	messageChan := make(chan []byte)
	
	// TODO: Register this channel with the notification service
	// This is just a placeholder implementation
	go func() {
		// Simulate sending message events every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				event := fmt.Sprintf("data: {\"type\":\"message\",\"message_id\":%d,\"sender_id\":\"user123\",\"content\":\"Hello there!\"}\n\n", time.Now().Unix())
				messageChan <- []byte(event)
			case <-r.Context().Done():
				close(messageChan)
				return
			}
		}
	}()

	// Stream events to client
	flusher, ok := w.(http.Flusher)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Streaming not supported")
		return
	}

	for {
		select {
		case <-r.Context().Done():
			return
		case msg, ok := <-messageChan:
			if !ok {
				return
			}
			_, err := w.Write(msg)
			if err != nil {
				return
			}
			flusher.Flush()
		}
	}
}

// NotificationEvents handles SSE for general notifications
func (h *NotificationHandler) NotificationEvents(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create notification channel
	notificationChan := make(chan []byte)
	
	// TODO: Register this channel with the notification service
	// This is just a placeholder implementation
	go func() {
		// Simulate sending notification events
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				notificationTypes := []string{"new_match", "new_like", "new_message"}
				notificationType := notificationTypes[time.Now().Unix()%3]
				event := fmt.Sprintf("data: {\"type\":\"%s\",\"user_id\":\"user456\",\"timestamp\":%d}\n\n", 
					notificationType, time.Now().Unix())
				notificationChan <- []byte(event)
			case <-r.Context().Done():
				close(notificationChan)
				return
			}
		}
	}()

	// Stream events to client
	flusher, ok := w.(http.Flusher)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Streaming not supported")
		return
	}

	for {
		select {
		case <-r.Context().Done():
			return
		case notification, ok := <-notificationChan:
			if !ok {
				return
			}
			_, err := w.Write(notification)
			if err != nil {
				return
			}
			flusher.Flush()
		}
	}
} 