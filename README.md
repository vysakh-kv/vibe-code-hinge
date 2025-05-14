# Vibe - A Hinge Clone

Vibe is a dating app inspired by Hinge, connecting people who match each other's vibe.

## Project Structure

This repository contains both the frontend and backend code:

- `/frontend` - Vue.js application built with Vite
- `/backend` - Go API server with Gorilla Mux

## Getting Started

### Prerequisites

- Node.js and npm
- Go 1.18+
- PostgreSQL
- Supabase account (for auth and storage)
- Docker and Docker Compose (optional, for containerized setup)

### Setup

#### Option 1: Manual Setup

1. Clone the repository:
```
git clone https://github.com/yourusername/vibe-code-hinge.git
cd vibe-code-hinge
```

2. Set up the backend:
```
cd backend
go mod download
cp .env.example .env  # Then edit .env with your credentials
```

3. Run database migrations:
```
go run cmd/migrate/main.go up
```

4. Set up the frontend:
```
cd ../frontend
npm install
cp .env.example .env  # Then edit .env with your credentials
```

#### Option 2: Using Makefile

```
# Set up both backend and frontend
make setup

# Run database migrations
make migrate-up
```

#### Option 3: Using Docker Compose

```
# Start all services
make docker-up

# Run migrations in Docker
make docker-migrate
```

#### Option 4: Using Render PostgreSQL Database

The project is configured to work with a Render hosted PostgreSQL database. To use the existing database:

```
# Set up environment with Render PostgreSQL credentials
make setup-env

# Then run migrations
make migrate-up
```

### Running the Application

#### Option 1: Manual Start

1. Start the backend server:
```
cd backend
go run cmd/api/main.go
```

2. Start the frontend development server:
```
cd frontend
npm run dev
```

#### Option 2: Using Makefile

```
# Start both servers
make start

# Or start them individually
make start-backend
make start-frontend
```

#### Option 3: Using Docker Compose

```
# Start all services
make docker-up

# View logs
make docker-logs

# Stop all services
make docker-down
```

## Features

- User authentication
- Profile creation and editing
- Discovering potential matches
- Swiping functionality
- Matching system
- Real-time messaging

## Technology Stack

- **Frontend**: Vue 3, Vue Router, Pinia, Vite
- **Backend**: Go, Gorilla Mux
- **Database**: PostgreSQL (hosted on Render)
- **Auth & Storage**: Supabase
- **Deployment**: Render (API + PostgreSQL)
- **Containerization**: Docker & Docker Compose

## Development Commands

All common development tasks are available through the Makefile:

```
# Show all available commands
make help
```

## Development Guidelines

When working on this project, please follow our development guidelines:

- General development guidelines are in the `llm_context.txt` file
- If you're using Cursor AI, please refer to the `CURSOR_GUIDELINES.md` file for additional guidelines

These guidelines help maintain consistency and quality across the codebase.

## Project Status

This is an MVP (Minimum Viable Product) version of the application.

## License

MIT
