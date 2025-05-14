package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vibe-code-hinge/backend/internal/models"
	"github.com/vibe-code-hinge/backend/internal/services"
)

// MessageHandler handles message-related routes
type MessageHandler struct {
	messageService *services.MessageService
}

// NewMessageHandler creates a new message handler
func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
	}
}

// GetMessages retrieves messages for a match
func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get match ID from path
	vars := mux.Vars(r)
	matchIDStr := vars["id"]
	matchID, err := strconv.ParseInt(matchIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid match ID")
		return
	}

	// Get pagination parameters
	limit := getIntQueryParam(r, "limit", 50)
	offset := getIntQueryParam(r, "offset", 0)

	// Get messages
	messages, err := h.messageService.GetMessages(r.Context(), userID, matchID, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

// CreateMessage sends a message in a match
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get match ID from path
	vars := mux.Vars(r)
	matchIDStr := vars["id"]
	matchID, err := strconv.ParseInt(matchIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid match ID")
		return
	}

	// Parse request body
	var input struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input
	if input.Message == "" {
		respondWithError(w, http.StatusBadRequest, "Message is required")
		return
	}

	// Create message input
	messageInput := models.MessageInput{
		MatchID: matchID,
		Message: input.Message,
	}

	// Send message
	message, err := h.messageService.SendMessage(r.Context(), userID, messageInput)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, message)
}

// Helper function to get a query parameter as an integer
func getIntQueryParam(r *http.Request, param string, defaultValue int) int {
	valueStr := r.URL.Query().Get(param)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(valueStr)
	if err != nil || value < 0 {
		return defaultValue
	}
	
	return value
}
