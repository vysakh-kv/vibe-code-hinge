package handlers

import (
	"net/http"

	"github.com/vibe-code-hinge/backend/internal/services"
)

// PromptHandler handles prompt-related routes
type PromptHandler struct {
	promptService *services.PromptService
}

// NewPromptHandler creates a new prompt handler
func NewPromptHandler(promptService *services.PromptService) *PromptHandler {
	return &PromptHandler{
		promptService: promptService,
	}
}

// GetDefaultPrompts handles the retrieval of default prompts
func (h *PromptHandler) GetDefaultPrompts(w http.ResponseWriter, r *http.Request) {
	// Call service to get default prompts
	prompts, err := h.promptService.GetDefaultPrompts(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, prompts)
} 