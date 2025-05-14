package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vibe-code-hinge/backend/internal/services"
)

// PreferenceHandler handles preference-related routes
type PreferenceHandler struct {
	preferenceService *services.PreferenceService
}

// NewPreferenceHandler creates a new preference handler
func NewPreferenceHandler(preferenceService *services.PreferenceService) *PreferenceHandler {
	return &PreferenceHandler{
		preferenceService: preferenceService,
	}
}

// GetPreferences handles the retrieval of user preferences
func (h *PreferenceHandler) GetPreferences(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Call service to get preferences
	preferences, err := h.preferenceService.GetPreferences(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, preferences)
}

// UpdatePreferences handles the update of user preferences
func (h *PreferenceHandler) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (would come from JWT middleware)
	// For demonstration, we'll use a query parameter for now
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	var preferences map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&preferences); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Call service to update preferences
	err := h.preferenceService.UpdatePreferences(r.Context(), userID, preferences)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
} 