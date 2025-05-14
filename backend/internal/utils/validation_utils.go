package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

// ValidationError holds information about a validation error
type ValidationError struct {
	Field   string
	Message string
}

// ValidationErrors holds multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError
}

// Add adds a validation error
func (ve *ValidationErrors) Add(field, message string) {
	ve.Errors = append(ve.Errors, ValidationError{Field: field, Message: message})
}

// HasErrors returns true if there are validation errors
func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.Errors) > 0
}

// Error implements the error interface
func (ve *ValidationErrors) Error() string {
	if len(ve.Errors) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Validation errors:\n")
	for _, err := range ve.Errors {
		sb.WriteString(fmt.Sprintf("- %s: %s\n", err.Field, err.Message))
	}
	return sb.String()
}

// ValidateEmail validates an email address
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email format: %w", err)
	}

	return nil
}

// ValidatePassword validates a password
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	var missing []string
	if !hasUpper {
		missing = append(missing, "uppercase letter")
	}
	if !hasLower {
		missing = append(missing, "lowercase letter")
	}
	if !hasDigit {
		missing = append(missing, "digit")
	}
	if !hasSpecial {
		missing = append(missing, "special character")
	}

	if len(missing) > 0 {
		return fmt.Errorf("password must contain at least one %s", strings.Join(missing, ", "))
	}

	return nil
}

// ValidateUsername validates a username
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if len(username) > 30 {
		return fmt.Errorf("username cannot be longer than 30 characters")
	}

	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validUsername.MatchString(username) {
		return fmt.Errorf("username can only contain letters, numbers, and underscores")
	}

	return nil
}

// ValidatePhoneNumber validates a phone number
func ValidatePhoneNumber(phone string) error {
	if phone == "" {
		return fmt.Errorf("phone number cannot be empty")
	}

	// Remove spaces, dashes, parentheses
	clean := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r
		}
		return -1
	}, phone)

	// Check if it starts with + and has 10-15 digits
	if !regexp.MustCompile(`^\+?[0-9]{10,15}$`).MatchString(clean) {
		return fmt.Errorf("invalid phone number format")
	}

	return nil
}

// ValidateNotEmpty validates that a string is not empty
func ValidateNotEmpty(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s cannot be empty", field)
	}
	return nil
}

// ValidateLength validates that a string has a length within the specified range
func ValidateLength(field, value string, min, max int) error {
	if len(value) < min {
		return fmt.Errorf("%s must be at least %d characters long", field, min)
	}
	if len(value) > max {
		return fmt.Errorf("%s cannot be longer than %d characters", field, max)
	}
	return nil
}

// ValidateURL validates a URL
func ValidateURL(url string) error {
	if url == "" {
		return nil // Empty URLs are allowed
	}

	// Basic URL validation - can be expanded as needed
	validURL := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(?:/[a-zA-Z0-9\-\._~:/?#[\]@!$&'()*+,;=]*)?$`)
	if !validURL.MatchString(url) {
		return fmt.Errorf("invalid URL format")
	}

	return nil
}

// ValidateAge validates that an age is within a reasonable range
func ValidateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("must be at least 18 years old")
	}
	if age > 120 {
		return fmt.Errorf("age cannot be greater than 120")
	}
	return nil
} 