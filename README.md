# BudgetCLI

A command-line interface (CLI) application for budget management.

---
## Overview

BudgetCLI is a tool designed to help users manage their budgets directly from the terminal. It uses SQLite for data storage and Goose for database migrations.

---
## Requirements

- **Go**: v1.25 or higher
- **SQLite**: Ensure you have SQLite installed if you want to inspect the database manually.
- **Goose**: Used for database migrations.

---
## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd budgetcli
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Configure environment variables**:
   Create a `.env` file in the root directory (refer to [Environment Variables](#environment-variables)).

4. **Run migrations**:
   ```bash
   # Using goose (ensure goose is installed)
   goose -dir database/migrations sqlite3 database/app.db up
   ```

---
## Run Commands

To run the application:

```bash
go run app/cli/main.go
```

To build the application:

```bash
go build -o budgetcli app/cli/main.go
./budgetcli
```

---
## Scripts

Currently, there are no dedicated script files (like `Makefile` or `scripts/`). Use the standard Go CLI commands.

- **Build**: `go build -o budgetcli app/cli/main.go`
- **Run**: `go run app/cli/main.go`
- **Migrate Up**: `goose -dir database/migrations sqlite3 database/app.db up`

---
## Environment Variables

The application uses the following environment variables (typically stored in a `.env` file):

| Variable | Description | Example Value |
|----------|-------------|---------------|
| `GOOSE_DRIVER` | Database driver for Goose | `sqlite3` |
| `GOOSE_DBSTRING` | Path to the SQLite database file | `database/app.db` |
| `GOOSE_MIGRATION_DIR` | Path to the migrations directory | `database/migrations` |

---
## Tests

TODO: No tests are currently implemented. 

To run tests (when added):
```bash
go test ./...
```

---
## Project Structure

```text
.
├── app/
│   └── cli/
│       └── main.go       # Application entry point
├── database/
│   ├── app.db            # SQLite database (git-ignored)
│   └── migrations/       # SQL migration files
├── internal/             # Internal packages (currently empty)
├── .env                  # Environment variables (git-ignored)
├── go.mod                # Go module definition
└── go.sum                # Go module checksums
```

---
## License

TODO: Add license information.
