package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Connect to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Check if connection is alive
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Connected to database successfully")

	// Query for table names
	rows, err := db.Query(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public'
		ORDER BY table_name
	`)
	if err != nil {
		log.Fatalf("Error querying tables: %v", err)
	}
	defer rows.Close()

	// Print table names
	fmt.Println("\nTables in the database:")
	fmt.Println("----------------------")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Println(tableName)
	}

	// Check for schema_migrations table specifically
	var exists bool
	err = db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 
			FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = 'schema_migrations'
		)
	`).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking for schema_migrations table: %v", err)
	}

	if exists {
		fmt.Println("\nMigration table exists, checking version...")
		var version int
		var dirty bool
		err = db.QueryRow(`SELECT version, dirty FROM schema_migrations`).Scan(&version, &dirty)
		if err != nil {
			log.Fatalf("Error querying migration version: %v", err)
		}
		fmt.Printf("Current migration version: %d, Dirty: %t\n", version, dirty)
	} else {
		fmt.Println("\nMigration table does not exist!")
	}
} 