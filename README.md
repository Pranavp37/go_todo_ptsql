# Go Todo API

A simple REST API for managing todos with user authentication, built with Go, Echo, and PostgreSQL.

## Tech Stack

- **Go** - Programming language
- **Echo** - Web framework
- **PostgreSQL** - Database
- **pgx** - PostgreSQL driver
- **JWT** - Authentication
- **Goose** - Database migrations

## Project Structure

```
├── cmd/api/          # Application entry point
├── internal/
│   ├── config/       # Configuration loading
│   ├── database/     # Database connection
│   ├── handlers/     # HTTP handlers
│   ├── middleware/   # JWT auth & other middleware
│   ├── models/       # Data models
│   ├── repository/   # Database queries
│   └── utiles/       # Helper functions
├── migrations/       # SQL migration files
└── makefile          # Build & migration commands
```

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) (for migrations)

### Setup

1. **Clone the repo**
   ```bash
   git clone https://github.com/Pranavp37/go_todo_ptsql.git
   cd go_todo_ptsql
   ```

2. **Create `.env` file**
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/todo_db?sslmode=disable
   GOOSE_DRIVER=postgres
   GOOSE_DBSTRING=postgres://user:password@localhost:5432/todo_db?sslmode=disable
   GOOSE_MIGRATION_DIR=migrations
   JWT_SECRET=your-secret-key
   ```

3. **Run migrations**
   ```bash
   make goose-up
   ```

4. **Start the server**
   ```bash
   go run cmd/api/main.go
   ```

   Server runs on `http://localhost:8080`

## API Endpoints

### Public Routes

| Method | Endpoint  | Description       |
|--------|-----------|-------------------|
| GET    | `/`       | Health check      |
| POST   | `/create` | Register new user |
| POST   | `/login`  | Login user        |

### Protected Routes (Requires JWT)

| Method | Endpoint            | Description       |
|--------|---------------------|-------------------|
| GET    | `/user/:id`         | Get user by ID    |
| POST   | `/create-todo`      | Create a new todo |
| GET    | `/todos`            | Get all todos     |
| PUT    | `/update-todo/:id`  | Update a todo     |
| DELETE | `/delete-todo/:id`  | Delete a todo     |

## Makefile Commands

```bash
make goose-up       # Run all pending migrations
make goose-down     # Rollback last migration
make goose-status   # Show migration status
make goose-reset    # Reset all migrations
make goose-create name=migration_name  # Create new migration
```

## License

MIT
