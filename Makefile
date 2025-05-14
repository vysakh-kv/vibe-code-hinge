.PHONY: setup start start-backend start-frontend build migrate-up migrate-down test help docker-up docker-down docker-logs docker-migrate setup-env update-context backend-logs clean-logs start-backend-debug

# Variables
LOG_LEVEL ?= info

# Default target
all: help

# Setup the project
setup: setup-backend setup-frontend

setup-backend:
	@echo "Setting up backend..."
	cd backend && go mod download
	cp backend/.env.example backend/.env
	@echo "Backend setup complete! Remember to update the .env file with your credentials."

setup-frontend:
	@echo "Setting up frontend..."
	cd frontend && npm install
	cp frontend/.env.example frontend/.env
	@echo "Frontend setup complete! Remember to update the .env file with your credentials."

# Setup environment with Render PostgreSQL credentials
setup-env:
	@echo "Setting up Render PostgreSQL environment variables..."
	@echo "# Database" > backend/.env
	@echo "# Option 1: Full connection string" >> backend/.env
	@echo "DATABASE_URL=postgresql://vibe_code_hinge_user:b9NCFB2Xdw70qlDjlucV3fuuJ30idGdz@dpg-d0i7eb24d50c73b63cfg-a.oregon-postgres.render.com/vibe_code_hinge?sslmode=require" >> backend/.env
	@echo "" >> backend/.env
	@echo "# Option 2: Individual connection parameters" >> backend/.env
	@echo "DB_USER=vibe_code_hinge_user" >> backend/.env
	@echo "DB_PASSWORD=b9NCFB2Xdw70qlDjlucV3fuuJ30idGdz" >> backend/.env
	@echo "DB_HOST=dpg-d0i7eb24d50c73b63cfg-a.oregon-postgres.render.com" >> backend/.env
	@echo "DB_NAME=vibe_code_hinge" >> backend/.env
	@echo "DB_PORT=5432" >> backend/.env
	@echo "DB_SSL_MODE=require" >> backend/.env
	@echo "" >> backend/.env
	@echo "# Server" >> backend/.env
	@echo "PORT=8080" >> backend/.env
	@echo "" >> backend/.env
	@echo "# Supabase" >> backend/.env
	@echo "SUPABASE_URL=https://asydpwsbayeejyhkxeih.supabase.co" >> backend/.env
	@echo "SUPABASE_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFzeWRwd3NiYXllZWp5aGt4ZWloIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDcyMjA3OTcsImV4cCI6MjA2Mjc5Njc5N30.dcUCQM2poNFWdlEaQTjNPKNq2eNt-hsq-s1Rx_gUZsM" >> backend/.env
	@echo "" >> backend/.env
	@echo "# JWT" >> backend/.env
	@echo "JWT_SECRET=your_jwt_secret_here" >> backend/.env
	@echo "JWT_EXPIRY=24h" >> backend/.env
	@echo "Environment file created successfully with Render PostgreSQL credentials."
	@echo "Setting up frontend environment variables..."
	@echo "VITE_API_URL=http://localhost:8080/api/v1" > frontend/.env
	@echo "VITE_SUPABASE_URL=https://asydpwsbayeejyhkxeih.supabase.co" >> frontend/.env
	@echo "VITE_SUPABASE_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFzeWRwd3NiYXllZWp5aGt4ZWloIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDcyMjA3OTcsImV4cCI6MjA2Mjc5Njc5N30.dcUCQM2poNFWdlEaQTjNPKNq2eNt-hsq-s1Rx_gUZsM" >> frontend/.env
	@echo "Frontend environment file created successfully."

# Remind to update LLM context file
update-context:
	@echo "\033[1;33m┌────────────────────────────────────────────────┐"
	@echo "│ REMINDER: Update the LLM Context File              │"
	@echo "├────────────────────────────────────────────────┤"
	@echo "│                                                │"
	@echo "│ If you've made significant changes, please     │"
	@echo "│ update the llm_context.txt file to help AI     │"
	@echo "│ assistants and other developers understand     │"
	@echo "│ the current state of the project.              │"
	@echo "│                                                │"
	@echo "│ Opening llm_context.txt for editing...         │"
	@echo "└────────────────────────────────────────────────┘\033[0m"
	@sleep 2
	@${EDITOR} llm_context.txt || echo "Could not open editor. Please edit llm_context.txt manually."

# Start development servers
start: start-backend start-frontend

start-backend:
	@echo "Setting up logs directory..."
	@mkdir -p logs
	@echo "Starting backend server with LOG_LEVEL=$(LOG_LEVEL)..."
	@cd backend && LOG_LEVEL=$(LOG_LEVEL) go run cmd/api/main.go 2>&1 | tee ../logs/backend.log &
	@echo "Backend server started in background. Logs available in logs/backend.log"
	@echo "Run 'make backend-logs' to follow logs in real-time"

start-frontend:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

# Build for production
build: build-backend build-frontend update-context

build-backend:
	@echo "Building backend..."
	cd backend && go build -o bin/api cmd/api/main.go
	@echo "Backend build complete!"

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm run build
	@echo "Frontend build complete!"

# Database migrations
migrate-up:
	@echo "Running database migrations up..."
	cd backend && go run cmd/migrate/main.go up

migrate-down:
	@echo "Running database migrations down..."
	cd backend && go run cmd/migrate/main.go down

# Run tests
test: test-backend test-frontend

test-backend:
	@echo "Running backend tests..."
	cd backend && go test ./...

test-frontend:
	@echo "Running frontend tests..."
	cd frontend && npm test

# Docker commands
docker-up:
	@echo "Starting all services with Docker Compose..."
	docker-compose up -d

docker-down:
	@echo "Stopping all services with Docker Compose..."
	docker-compose down

docker-logs:
	@echo "Showing logs from all services..."
	docker-compose logs -f

docker-migrate:
	@echo "Running database migrations in Docker..."
	docker-compose exec backend go run cmd/migrate/main.go up

# View backend logs
backend-logs:
	@echo "Viewing backend logs in real-time..."
	@tail -f logs/backend.log

# Show help
help:
	@echo "Available commands:"
	@echo "  make setup                - Set up the project (both backend and frontend)"
	@echo "  make setup-backend        - Set up the backend only"
	@echo "  make setup-frontend       - Set up the frontend only"
	@echo "  make setup-env            - Set up environment with Render PostgreSQL credentials"
	@echo "  make update-context       - Reminder to update the LLM context file after changes"
	@echo "  make start                - Start both backend and frontend development servers"
	@echo "  make start-backend        - Start the backend server only"
	@echo "  make start-backend-debug  - Start the backend server with DEBUG logging"
	@echo "  make start-frontend       - Start the frontend server only"
	@echo "  make build                - Build both backend and frontend for production"
	@echo "  make build-backend        - Build the backend only"
	@echo "  make build-frontend       - Build the frontend only"
	@echo "  make migrate-up           - Run database migrations up"
	@echo "  make migrate-down         - Run database migrations down"
	@echo "  make test                 - Run all tests"
	@echo "  make test-backend         - Run backend tests only"
	@echo "  make test-frontend        - Run frontend tests only"
	@echo "  make docker-up            - Start all services with Docker Compose"
	@echo "  make docker-down          - Stop all services with Docker Compose"
	@echo "  make docker-logs          - Show logs from all services"
	@echo "  make docker-migrate       - Run database migrations in Docker"
	@echo "  make backend-logs         - View backend logs in real-time"
	@echo "  make clean-logs           - Remove all log files"
	@echo "  make help                 - Show this help message"

# Clean log files
clean-logs:
	@echo "Cleaning log files..."
	@rm -rf logs
	@echo "Log files removed."

# Start backend with debug logging
start-backend-debug:
	@make start-backend LOG_LEVEL=debug
	@echo "Backend started with DEBUG log level" 