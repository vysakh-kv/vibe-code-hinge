package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vibe-code-hinge/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles user-related business logic
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new user service
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// Register creates a new user account
func (s *UserService) Register(ctx context.Context, input models.UserInput) (*models.AuthResponse, error) {
	// Check if user already exists
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", input.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	var user models.User
	now := time.Now()

	err = s.db.QueryRowContext(
		ctx,
		"INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, email, created_at, updated_at",
		input.Email, string(hashedPassword), now, now,
	).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Generate JWT token (implementation omitted)
	token := "dummy-token" // Replace with actual JWT implementation

	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// Login authenticates a user
func (s *UserService) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	var user models.User
	var hashedPassword string

	// Get user by email
	err := s.db.QueryRowContext(
		ctx,
		"SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1",
		input.Email,
	).Scan(&user.ID, &user.Email, &hashedPassword, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token (implementation omitted)
	token := "dummy-token" // Replace with actual JWT implementation

	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User

	err := s.db.QueryRowContext(
		ctx,
		"SELECT id, email, created_at, updated_at FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
