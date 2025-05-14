package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Constants for token expiration
const (
	// TokenExpiration is the default expiration time for tokens
	TokenExpiration = 24 * time.Hour
	
	// RefreshTokenExpiration is the default expiration time for refresh tokens
	RefreshTokenExpiration = 7 * 24 * time.Hour
)

// TokenClaims holds the claims for a token
type TokenClaims struct {
	UserID    int64  `json:"uid,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plain text password
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken generates a signed token with claims
func GenerateToken(claims TokenClaims, secret string) (string, error) {
	// Set standard claims if not set
	if claims.ExpiresAt == 0 {
		claims.ExpiresAt = time.Now().Add(TokenExpiration).Unix()
	}
	if claims.IssuedAt == 0 {
		claims.IssuedAt = time.Now().Unix()
	}
	if claims.NotBefore == 0 {
		claims.NotBefore = time.Now().Unix()
	}

	// Create header
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	// Encode header and claims
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal header: %w", err)
	}
	
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("failed to marshal claims: %w", err)
	}
	
	headerBase64 := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsBase64 := base64.RawURLEncoding.EncodeToString(claimsJSON)
	
	// Create signature
	signingString := headerBase64 + "." + claimsBase64
	signature := generateSignature(signingString, secret)
	
	// Return the token
	return signingString + "." + signature, nil
}

// ValidateToken validates a token and returns the claims
func ValidateToken(tokenString, secret string) (*TokenClaims, error) {
	// Split the token
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}
	
	// Verify signature
	signingString := parts[0] + "." + parts[1]
	signature := generateSignature(signingString, secret)
	
	if signature != parts[2] {
		return nil, fmt.Errorf("invalid token signature")
	}
	
	// Decode claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode claims: %w", err)
	}
	
	var claims TokenClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims: %w", err)
	}
	
	// Validate claims
	now := time.Now().Unix()
	if claims.ExpiresAt != 0 && now > claims.ExpiresAt {
		return nil, fmt.Errorf("token has expired")
	}
	if claims.NotBefore != 0 && now < claims.NotBefore {
		return nil, fmt.Errorf("token not valid yet")
	}
	
	return &claims, nil
}

// GenerateUserToken generates a token for a user
func GenerateUserToken(userID int64, email, username, secret string) (string, error) {
	claims := TokenClaims{
		UserID:    userID,
		Email:     email,
		Username:  username,
		ExpiresAt: time.Now().Add(TokenExpiration).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
	}
	
	return GenerateToken(claims, secret)
}

// GenerateRefreshToken generates a refresh token
func GenerateRefreshToken(userID int64, secret string) (string, error) {
	claims := TokenClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: time.Now().Add(RefreshTokenExpiration).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
	}
	
	return GenerateToken(claims, secret)
}

// ValidateRefreshToken validates a refresh token
func ValidateRefreshToken(tokenString, secret string) (int64, error) {
	claims, err := ValidateToken(tokenString, secret)
	if err != nil {
		return 0, err
	}
	
	if claims.Subject == "" {
		return 0, fmt.Errorf("invalid subject in refresh token")
	}
	
	var userID int64
	_, err = fmt.Sscanf(claims.Subject, "%d", &userID)
	if err != nil {
		return 0, fmt.Errorf("invalid subject in refresh token: %w", err)
	}
	
	return userID, nil
}

// generateSignature generates a signature for a signing string
func generateSignature(signingString, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signingString))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// GenerateRandomBytes generates n random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	
	return b, nil
}

// GenerateRandomString generates a random string of length n
func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	
	return base64.URLEncoding.EncodeToString(b)[:n], nil
}

// GenerateAPIKey generates a random API key
func GenerateAPIKey() (string, error) {
	return GenerateRandomString(32)
}

// GenerateSessionID generates a random session ID
func GenerateSessionID() (string, error) {
	return GenerateRandomString(64)
}

// SanitizeSensitiveData removes sensitive data from a map
func SanitizeSensitiveData(data map[string]interface{}) map[string]interface{} {
	sensitiveFields := []string{"password", "password_hash", "api_key", "secret", "token", "access_token", "refresh_token"}
	result := make(map[string]interface{})
	
	for k, v := range data {
		isSensitive := false
		for _, field := range sensitiveFields {
			if k == field {
				isSensitive = true
				break
			}
		}
		
		if isSensitive {
			result[k] = "********"
		} else {
			result[k] = v
		}
	}
	
	return result
} 