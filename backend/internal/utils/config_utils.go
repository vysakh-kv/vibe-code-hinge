package utils

import (
	"os"
	"strconv"
	"time"
)

// Config is a utility struct to handle environment variables
type Config struct{}

// NewConfig creates a new Config
func NewConfig() *Config {
	return &Config{}
}

// GetEnv gets an environment variable with a default value
func (c *Config) GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnvInt gets an integer environment variable with a default value
func (c *Config) GetEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// GetEnvBool gets a boolean environment variable with a default value
func (c *Config) GetEnvBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// GetEnvDuration gets a duration environment variable with a default value
func (c *Config) GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// GetDatabaseURL constructs a database URL from individual parameters or returns the DATABASE_URL environment variable
func (c *Config) GetDatabaseURL() string {
	// If DATABASE_URL is set, use it
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		return dbURL
	}
	
	// Otherwise, build it from individual parameters
	dbUser := c.GetEnv("DB_USER", "")
	dbPassword := c.GetEnv("DB_PASSWORD", "")
	dbHost := c.GetEnv("DB_HOST", "localhost")
	dbName := c.GetEnv("DB_NAME", "")
	dbPort := c.GetEnv("DB_PORT", "5432")
	dbSSLMode := c.GetEnv("DB_SSL_MODE", "require")
	
	// Check if required parameters are set
	if dbUser == "" || dbPassword == "" || dbName == "" {
		return ""
	}
	
	// Construct the URL
	return "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
} 