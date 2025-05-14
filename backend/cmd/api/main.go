package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vibe-code-hinge/backend/internal/routes"
)

func main() {
	// Configure structured logging
	logLevel := getLogLevel()
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	slog.Info("Starting application", "log_level", logLevel.String())

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using environment variables")
	}

	// Database connection
	var dbURL string
	dbURL = os.Getenv("DATABASE_URL")

	// If DATABASE_URL is not set, build it from individual parameters
	if dbURL == "" {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dbPort := os.Getenv("DB_PORT")
		dbSSLMode := os.Getenv("DB_SSL_MODE")

		if dbUser == "" || dbPassword == "" || dbHost == "" || dbName == "" {
			slog.Error("Database connection parameters are not set properly")
			os.Exit(1)
		}

		if dbPort == "" {
			dbPort = "5432" // Default PostgreSQL port
		}

		if dbSSLMode == "" {
			dbSSLMode = "require" // Default to require SSL
		}

		dbURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	}

	if dbURL == "" {
		slog.Error("DATABASE_URL or individual database parameters must be set")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Check database connection
	if err := db.Ping(); err != nil {
		slog.Error("Error pinging database", "error", err)
		os.Exit(1)
	}
	slog.Info("Connected to database successfully")

	// Initialize router
	router := mux.NewRouter()
	
	// Middleware for CORS and logging
	router.Use(corsMiddleware)
	router.Use(loggingMiddleware)

	// API routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.SetupRoutes(apiRouter, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info("Server starting", "port", port)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}

// getLogLevel returns the appropriate slog.Level based on the LOG_LEVEL environment variable
func getLogLevel() slog.Level {
	levelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	
	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo // Default to info level
	}
}

// corsMiddleware handles CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response wrapper to capture the status code
		wrapper := newResponseWriter(w)
		
		// Call the next handler
		next.ServeHTTP(wrapper, r)
		
		// Log the request details
		slog.Info("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapper.status,
			"duration_ms", time.Since(start).Milliseconds(),
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)
	})
}

// responseWriter is a wrapper for http.ResponseWriter that captures the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

// newResponseWriter creates a new responseWriter
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
} 
