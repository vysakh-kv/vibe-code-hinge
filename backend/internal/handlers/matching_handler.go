package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
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
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	var input models.SwipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input
	if input.ProfileID == "" {
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
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
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

// LikeProfile handles the action of liking a profile
func (h *MatchingHandler) LikeProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get profile ID from URL path
	vars := mux.Vars(r)
	profileID := vars["id"]
	if profileID == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	// Call service to create a like
	match, err := h.matchingService.CreateSwipe(r.Context(), userID, profileID, true)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If this was a match, return it, otherwise return success
	if match != nil {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("It's a match!", match))
	} else {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("Like recorded", nil))
	}
}

// SkipProfile handles the action of skipping a profile
func (h *MatchingHandler) SkipProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get profile ID from URL path
	vars := mux.Vars(r)
	profileID := vars["id"]
	if profileID == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	// Call service to create a skip
	_, err := h.matchingService.CreateSwipe(r.Context(), userID, profileID, false)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("Skip recorded", nil))
}

// SendRose handles the action of sending a rose to a profile
func (h *MatchingHandler) SendRose(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Get profile ID from URL path
	vars := mux.Vars(r)
	profileID := vars["id"]
	if profileID == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	// TODO: Implement sending a rose (premium feature)
	// For now, just record it as a special type of like
	match, err := h.matchingService.CreateSwipe(r.Context(), userID, profileID, true)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If this was a match, return it, otherwise return success
	if match != nil {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("It's a match! (Rose sent)", match))
	} else {
		respondWithJSON(w, http.StatusOK, models.NewSuccessResponse("Rose sent", nil))
	}
}

// GetLikes handles the retrieval of profiles that liked the user
func (h *MatchingHandler) GetLikes(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// TODO: Implement getting likes from the service
	// For now, return placeholder data
	likes := []map[string]interface{}{
		{
			"id":        1,
			"name":      "Sample User 1",
			"age":       28,
			"blurred":   true, // Requires premium to see clearly
			"timestamp": "2023-08-15T12:34:56Z",
		},
		{
			"id":        2,
			"name":      "Sample User 2",
			"age":       31,
			"blurred":   true,
			"timestamp": "2023-08-14T10:12:23Z",
		},
	}

	respondWithJSON(w, http.StatusOK, likes)
}
