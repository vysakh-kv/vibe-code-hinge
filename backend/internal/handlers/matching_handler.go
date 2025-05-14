package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vibe-code-hinge/backend/internal/models"
	"github.com/vibe-code-hinge/backend/internal/services"
)

// MatchingHandler handles matching-related routes
type MatchingHandler struct {
	matchingService *services.MatchingService
}

// NewMatchingHandler creates a new matching handler
func NewMatchingHandler(matchingService *services.MatchingService) *MatchingHandler {
	return &MatchingHandler{
		matchingService: matchingService,
	}
}

// GetDiscoverProfiles handles the retrieval of profiles for discovery
func (h *MatchingHandler) GetDiscoverProfiles(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Get limit from query params (default to 10)
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Call service to get discover profiles
	profiles, err := h.matchingService.GetDiscoverProfiles(r.Context(), userID, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, profiles)
}

// CreateSwipe handles the creation of a swipe
func (h *MatchingHandler) CreateSwipe(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var input models.SwipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input
	if input.ProfileID <= 0 {
		respondWithError(w, http.StatusBadRequest, "Profile ID is required")
		return
	}

	// Call service to create swipe
	match, err := h.matchingService.CreateSwipe(r.Context(), userID, input.ProfileID, input.IsLike)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If this was a match, return it, otherwise return success
	if match != nil {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("It's a match!", match))
	} else {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("Swipe recorded", nil))
	}
}

// GetMatches handles the retrieval of a user's matches
func (h *MatchingHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Call service to get matches
	matches, err := h.matchingService.GetMatches(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, matches)
}
