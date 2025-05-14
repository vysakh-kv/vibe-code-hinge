package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// DB is a wrapper around sql.DB that provides additional functionality
type DB struct {
	*sql.DB
}

// NewDB creates a new DB instance
func NewDB(db *sql.DB) *DB {
	return &DB{DB: db}
}

// Transactional executes a function within a database transaction
// If the function returns an error, the transaction is rolled back
// Otherwise, the transaction is committed
func Transactional(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	defer func() {
		if p := recover(); p != nil {
			// Rollback on panic
			tx.Rollback()
			panic(p) // Re-throw panic
		} else if err != nil {
			// Rollback on error
			tx.Rollback()
		} else {
			// Commit if all went well
			err = tx.Commit()
		}
	}()
	
	err = fn(tx)
	return err
}

// ExecuteWithTimeout executes a function with a timeout
func ExecuteWithTimeout(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	
	return fn(ctx)
}

// ExecuteQuery is a helper function to execute a query with parameters
func ExecuteQuery(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, args...)
}

// ExecuteQueryRow is a helper function to execute a query with parameters and return a single row
func ExecuteQueryRow(ctx context.Context, db *sql.DB, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}

// ExecuteUpdate is a helper function to execute a non-query SQL statement
func ExecuteUpdate(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.ExecContext(ctx, query, args...)
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	URL             string
}

// DefaultDatabaseConfig creates a default database configuration
func DefaultDatabaseConfig(url string) DatabaseConfig {
	return DatabaseConfig{
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxLifetime: time.Hour,
		ConnMaxIdleTime: 30 * time.Minute,
		URL:             url,
	}
}

// InitDatabase initializes a database connection with the given configuration
func InitDatabase(config DatabaseConfig) (*sql.DB, error) {
	if config.URL == "" {
		return nil, fmt.Errorf("database URL is required")
	}

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pool settings
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return db, nil
}

// CloseDB closes a database connection gracefully
func CloseDB(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v\n", err)
			return
		}
		log.Println("Database connection closed successfully")
	}
}

// Transaction executes the given function within a database transaction
func Transaction(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		// Error occurred, rollback
		if rbErr := tx.Rollback(); rbErr != nil {
			// Something really bad happened
			return fmt.Errorf("error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// NullStringToString converts sql.NullString to string safely
func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

// NullIntToInt64 converts sql.NullInt64 to int64 safely
func NullIntToInt64(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}
	return 0
}

// NullFloatToFloat64 converts sql.NullFloat64 to float64 safely
func NullFloatToFloat64(f sql.NullFloat64) float64 {
	if f.Valid {
		return f.Float64
	}
	return 0.0
}

// NullTimeToTime converts sql.NullTime to time.Time safely
func NullTimeToTime(t sql.NullTime) time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

// StringToNullString converts a string to sql.NullString
func StringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

// Int64ToNullInt64 converts an int64 to sql.NullInt64
func Int64ToNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: i, Valid: true}
}

// Float64ToNullFloat64 converts a float64 to sql.NullFloat64
func Float64ToNullFloat64(f float64) sql.NullFloat64 {
	if f == 0.0 {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: f, Valid: true}
}

// TimeToNullTime converts a time.Time to sql.NullTime
func TimeToNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: t, Valid: true}
} 