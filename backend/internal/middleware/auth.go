package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/vibe-code-hinge/backend/internal/models"
	"github.com/vibe-code-hinge/backend/internal/services"
)

// AuthKey is the context key for the authenticated user
type AuthKey string

// UserContextKey is the key for the user ID in the context
const UserContextKey AuthKey = "user"

// Auth middleware for authenticating requests
func Auth(userService *services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized: No Authorization header provided", http.StatusUnauthorized)
				return
			}

			// Extract token from "Bearer <token>"
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				http.Error(w, "Unauthorized: Invalid Authorization header format", http.StatusUnauthorized)
				return
			}
			token := strings.TrimSpace(splitToken[1])

			// Verify token with Supabase
			user, err := userService.VerifyToken(r.Context(), token)
			if err != nil {
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}

			// Add user to request context
			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserFromContext extracts the user ID from the request context
func GetUserFromContext(ctx context.Context) (int64, bool) {
	user, ok := ctx.Value(UserContextKey).(*models.User)
	if !ok || user == nil {
		return 0, false
	}
	return user.ID, true
}
