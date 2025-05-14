package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vibe-code-hinge/backend/internal/handlers"
	"github.com/vibe-code-hinge/backend/internal/services"
)

// SetupRoutes configures all the routes for the API
func SetupRoutes(router *mux.Router, db *sql.DB) {
	// Create services
	userService := services.NewUserService(db)
	profileService := services.NewProfileService(db)
	matchingService := services.NewMatchingService(db)
	messageService := services.NewMessageService(db)

	// Create handlers
	authHandler := handlers.NewAuthHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)
	matchingHandler := handlers.NewMatchingHandler(matchingService)
	messageHandler := handlers.NewMessageHandler(messageService)

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Auth routes
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Profile routes
	router.HandleFunc("/profiles/{id}", profileHandler.GetProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}", profileHandler.UpdateProfile).Methods("PUT")

	// Matching routes
	router.HandleFunc("/profiles/discover", matchingHandler.GetDiscoverProfiles).Methods("GET")
	router.HandleFunc("/swipes", matchingHandler.CreateSwipe).Methods("POST")
	router.HandleFunc("/matches", matchingHandler.GetMatches).Methods("GET")

	// Messaging routes
	router.HandleFunc("/matches/{id}/messages", messageHandler.GetMessages).Methods("GET")
	router.HandleFunc("/matches/{id}/messages", messageHandler.CreateMessage).Methods("POST")
}
