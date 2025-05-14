package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vibe-code-hinge/backend/internal/models"
)

// respondWithJSON writes a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error encoding response")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError writes an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.NewErrorResponse(message))
}
