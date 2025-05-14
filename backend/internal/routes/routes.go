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
	preferenceService := services.NewPreferenceService(db)
	promptService := services.NewPromptService(db)
	feedService := services.NewFeedService(db)
	notificationService := services.NewNotificationService(db)

	// Create handlers
	authHandler := handlers.NewAuthHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)
	matchingHandler := handlers.NewMatchingHandler(matchingService)
	messageHandler := handlers.NewMessageHandler(messageService)
	preferenceHandler := handlers.NewPreferenceHandler(preferenceService)
	promptHandler := handlers.NewPromptHandler(promptService)
	feedHandler := handlers.NewFeedHandler(feedService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

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
	router.HandleFunc("/profiles/{id}/prompts", profileHandler.GetProfilePrompts).Methods("GET")
	router.HandleFunc("/profiles/{id}/prompts", profileHandler.UpdateProfilePrompts).Methods("PUT")

	// User Preferences routes
	router.HandleFunc("/preferences", preferenceHandler.GetPreferences).Methods("GET")
	router.HandleFunc("/preferences", preferenceHandler.UpdatePreferences).Methods("PUT")

	// Prompts routes
	router.HandleFunc("/prompts", promptHandler.GetDefaultPrompts).Methods("GET")

	// Feed and Discovery routes
	router.HandleFunc("/feed", feedHandler.GetFeed).Methods("GET")
	router.HandleFunc("/standouts", feedHandler.GetStandouts).Methods("GET")
	router.HandleFunc("/profiles/discover", matchingHandler.GetDiscoverProfiles).Methods("GET") // Keep for backward compatibility

	// Interaction routes
	router.HandleFunc("/profiles/{id}/like", matchingHandler.LikeProfile).Methods("POST")
	router.HandleFunc("/profiles/{id}/skip", matchingHandler.SkipProfile).Methods("POST")
	router.HandleFunc("/profiles/{id}/rose", matchingHandler.SendRose).Methods("POST")
	router.HandleFunc("/swipes", matchingHandler.CreateSwipe).Methods("POST")

	// Likes and Matches routes
	router.HandleFunc("/likes", matchingHandler.GetLikes).Methods("GET")
	router.HandleFunc("/matches", matchingHandler.GetMatches).Methods("GET")

	// Messaging routes
	router.HandleFunc("/matches/{id}/messages", messageHandler.GetMessages).Methods("GET")
	router.HandleFunc("/matches/{id}/messages", messageHandler.CreateMessage).Methods("POST")

	// Server-Sent Events (SSE) routes
	router.HandleFunc("/events/messages", notificationHandler.MessageEvents).Methods("GET")
	router.HandleFunc("/events/notifications", notificationHandler.NotificationEvents).Methods("GET")
}
