package services

import (
	"database/sql"
)

// BaseService provides common functionality for all services
type BaseService struct {
	db *sql.DB
}

// NewBaseService creates a new BaseService
func NewBaseService(db *sql.DB) BaseService {
	return BaseService{db: db}
}

// GetDB returns the database instance
func (s *BaseService) GetDB() *sql.DB {
	return s.db
} 