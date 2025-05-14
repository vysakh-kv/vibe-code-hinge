package handlers

import (
	"net/http"
	"strconv"

	"github.com/vibe-code-hinge/backend/internal/services"
)

// FeedHandler handles feed-related routes
type FeedHandler struct {
	feedService *services.FeedService
}

// NewFeedHandler creates a new feed handler
func NewFeedHandler(feedService *services.FeedService) *FeedHandler {
	return &FeedHandler{
		feedService: feedService,
	}
}

// GetFeed handles the retrieval of the main feed
func (h *FeedHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get limit and offset from query params
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	
	limit := 10 // Default limit
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	
	offset := 0 // Default offset
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Call service to get feed
	feed, err := h.feedService.GetFeed(r.Context(), userID, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feed)
}

// GetStandouts handles the retrieval of standout profiles
func (h *FeedHandler) GetStandouts(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get limit from query params
	limitStr := r.URL.Query().Get("limit")
	limit := 10 // Default limit
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Call service to get standouts
	standouts, err := h.feedService.GetStandouts(r.Context(), userID, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, standouts)
} 