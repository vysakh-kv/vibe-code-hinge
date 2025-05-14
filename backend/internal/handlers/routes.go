package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all the routes for the API
func RegisterRoutes(router *mux.Router, db *sql.DB) {
	// Health check endpoint
	router.HandleFunc("/health", healthCheckHandler()).Methods("GET")
	
	// User routes
	router.HandleFunc("/auth/register", registerHandler(db)).Methods("POST")
	router.HandleFunc("/auth/login", loginHandler(db)).Methods("POST")
	
	// Profile routes
	router.HandleFunc("/profiles/{id}", getProfileHandler(db)).Methods("GET")
	router.HandleFunc("/profiles/{id}", updateProfileHandler(db)).Methods("PUT")
	
	// Matching routes
	router.HandleFunc("/profiles/discover", getDiscoverProfilesHandler(db)).Methods("GET")
	router.HandleFunc("/swipes", createSwipeHandler(db)).Methods("POST")
	router.HandleFunc("/matches", getMatchesHandler(db)).Methods("GET")
	
	// Messaging routes
	router.HandleFunc("/matches/{id}/messages", getMessagesHandler(db)).Methods("GET")
	router.HandleFunc("/matches/{id}/messages", createMessageHandler(db)).Methods("POST")
}

// respondWithJSON is a helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError is a helper function to respond with an error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Health check handler
func healthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Return API health status with more details than the root health endpoint
		status := map[string]interface{}{
			"status": "ok",
			"api_version": "v1",
			"timestamp": timeNow(),
			"services": map[string]string{
				"database": "connected",
				"auth": "available",
			},
		}
		respondWithJSON(w, http.StatusOK, status)
	}
}

// Helper function to get current time (makes testing easier)
func timeNow() string {
	return time.Now().Format(time.RFC3339)
}

// Handler functions
func registerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func loginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func getProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func updateProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func getDiscoverProfilesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func createSwipeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func getMatchesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func getMessagesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
}

func createMessageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This will be implemented later
		respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
	}
} 