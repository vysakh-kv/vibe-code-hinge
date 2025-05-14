package main

import (
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Create a new migrate instance
	m, err := migrate.New(
		"file://migrations/postgres",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Migration failed to initialize: %v", err)
	}

	// Check command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [up|down|version|force <version>]")
	}

	// Perform the requested migration operation
	command := os.Args[1]
	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Migration up complete")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Migration down complete")

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}
		log.Printf("Current migration version: %d, Dirty: %v", version, dirty)
		
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run cmd/migrate/main.go force <version>")
		}
		
		version, err := strconv.ParseUint(os.Args[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		
		if err := m.Force(int(version)); err != nil {
			log.Fatalf("Failed to force version: %v", err)
		}
		log.Printf("Successfully forced migration version to %d", version)

	default:
		log.Fatal("Invalid command. Use 'up', 'down', 'version', or 'force <version>'")
	}
} 