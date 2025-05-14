package utils

import (
	"context"
	"time"
)

// ContextKey is a type for context keys
type ContextKey string

// Context keys
const (
	// UserIDKey is the key for the user ID in a context
	UserIDKey ContextKey = "user_id"
	
	// UsernameKey is the key for the username in a context
	UsernameKey ContextKey = "username"
	
	// EmailKey is the key for the email in a context
	EmailKey ContextKey = "email"
	
	// RequestIDKey is the key for the request ID in a context
	RequestIDKey ContextKey = "request_id"
	
	// StartTimeKey is the key for the start time in a context
	StartTimeKey ContextKey = "start_time"
	
	// ClientIPKey is the key for the client IP in a context
	ClientIPKey ContextKey = "client_ip"
	
	// UserAgentKey is the key for the user agent in a context
	UserAgentKey ContextKey = "user_agent"
)

// GetUserID gets the user ID from a context
func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

// SetUserID sets the user ID in a context
func SetUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUsername gets the username from a context
func GetUsername(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)
	return username, ok
}

// SetUsername sets the username in a context
func SetUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, UsernameKey, username)
}

// GetEmail gets the email from a context
func GetEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(EmailKey).(string)
	return email, ok
}

// SetEmail sets the email in a context
func SetEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, EmailKey, email)
}

// GetRequestID gets the request ID from a context
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(RequestIDKey).(string)
	return requestID, ok
}

// SetRequestID sets the request ID in a context
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// GetStartTime gets the start time from a context
func GetStartTime(ctx context.Context) (time.Time, bool) {
	startTime, ok := ctx.Value(StartTimeKey).(time.Time)
	return startTime, ok
}

// SetStartTime sets the start time in a context
func SetStartTime(ctx context.Context, startTime time.Time) context.Context {
	return context.WithValue(ctx, StartTimeKey, startTime)
}

// GetClientIP gets the client IP from a context
func GetClientIP(ctx context.Context) (string, bool) {
	clientIP, ok := ctx.Value(ClientIPKey).(string)
	return clientIP, ok
}

// SetClientIP sets the client IP in a context
func SetClientIP(ctx context.Context, clientIP string) context.Context {
	return context.WithValue(ctx, ClientIPKey, clientIP)
}

// GetUserAgent gets the user agent from a context
func GetUserAgent(ctx context.Context) (string, bool) {
	userAgent, ok := ctx.Value(UserAgentKey).(string)
	return userAgent, ok
}

// SetUserAgent sets the user agent in a context
func SetUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, UserAgentKey, userAgent)
}

// WithTimeout returns a copy of the parent context with a timeout
func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

// ContextWithValues creates a new context with multiple values
func ContextWithValues(ctx context.Context, values map[ContextKey]interface{}) context.Context {
	for k, v := range values {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}

// GetElapsedTime gets the elapsed time since the start time in a context
func GetElapsedTime(ctx context.Context) (time.Duration, bool) {
	startTime, ok := GetStartTime(ctx)
	if !ok {
		return 0, false
	}
	return time.Since(startTime), true
}

// GetContextString gets a string value from a context
func GetContextString(ctx context.Context, key ContextKey) (string, bool) {
	val, ok := ctx.Value(key).(string)
	return val, ok
}

// GetContextInt64 gets an int64 value from a context
func GetContextInt64(ctx context.Context, key ContextKey) (int64, bool) {
	val, ok := ctx.Value(key).(int64)
	return val, ok
}

// GetContextBool gets a bool value from a context
func GetContextBool(ctx context.Context, key ContextKey) (bool, bool) {
	val, ok := ctx.Value(key).(bool)
	return val, ok
}

// GetContextTime gets a time.Time value from a context
func GetContextTime(ctx context.Context, key ContextKey) (time.Time, bool) {
	val, ok := ctx.Value(key).(time.Time)
	return val, ok
} 