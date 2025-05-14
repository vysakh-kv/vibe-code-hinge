package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vibe-code-hinge/backend/internal/models"
	"github.com/vibe-code-hinge/backend/internal/services"
)

// ProfileHandler handles profile routes
type ProfileHandler struct {
	profileService *services.ProfileService
}

// NewProfileHandler creates a new profile handler
func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

// GetProfile handles the retrieval of a profile
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	// Call service to get profile
	profile, err := h.profileService.GetProfileByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, profile)
}

// UpdateProfile handles the update of a profile
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	// Get user ID from context (would come from JWT middleware)
	// For now, we'll just use the profile ID as the user ID for simplicity
	userID := id

	var input models.ProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input
	if input.Name == "" || input.DateOfBirth == "" || input.Gender == "" {
		respondWithError(w, http.StatusBadRequest, "Name, date of birth, and gender are required")
		return
	}

	// Call service to update profile
	profile, err := h.profileService.CreateOrUpdateProfile(r.Context(), userID, input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, profile)
}
