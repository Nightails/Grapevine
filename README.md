# 🍇 Grapevine

A command-line interface (CLI) application for budget management.

---
## 📝 Overview

Grapevine is a terminal-based tool designed to help users manage their budgets directly from the command line. It uses SQLite for data storage, Goose for database migrations, and SQLC for type-safe database queries.

---
## 🛠️ Requirements

- **Go**: v1.21 or higher (using toolchain v1.23.4)
- **SQLite**: Ensure you have SQLite installed if you want to inspect the database manually.
- **Goose**: Used for database migrations.
- **SQLC**: Used to generate type-safe Go code from SQL.

---
## ⚙️ Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Nightails/Grapevine.git
   cd grapevine
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Configure environment variables (Optional)**:
   Refer to [docs/environment_variables.md](docs/environment_variables.md) for details on available environment variables.

4. **Run migrations**:
   ```bash
   # Using goose with explicit flags
   goose -dir database/migrations sqlite3 database/app.db up

   # OR using environment variables (if configured)
   goose up
   ```

5. **Generate SQL code**:
   ```bash
   # Using sqlc (ensure sqlc is installed)
   sqlc generate
   ```

---
## 🖥️ Run Commands

To run the application:

```bash
go run main.go
```

To build the application:

```bash
go build -o grapevine main.go
./grapevine
```

---
## 📔 Scripts

The following scripts are available in the `scripts/` directory:

- **Manage Test Data**: Use `scripts/manage_test_data.sh` to populate or clear the database with test data.
  - Populate: `./scripts/manage_test_data.sh up`
  - Clear: `./scripts/manage_test_data.sh down`

Standard Go CLI commands:

- **Build**: `go build -o grapevine main.go`
- **Run**: `go run main.go`
- **Migrate Up**: `goose -dir database/migrations sqlite3 database/app.db up` or simply `goose up` (if env vars are set)
- **SQL Generate**: `sqlc generate`

---
## 📄 Documentation

For detailed information about the database schema and environment variables, please refer to:
- [docs/database.md](docs/database.md)
- [docs/environment_variables.md](docs/environment_variables.md)

---
## 🪧 Tests

Automated Go tests are included in the project.

To run Go tests:
```bash
go test ./...
```

To populate test data:
```bash
./scripts/manage_test_data.sh up
```

---
## 📁 Project Structure

```text
.
├── database/
│   ├── app.db            # SQLite database (git-ignored)
│   ├── migrations/       # SQL migration files
│   ├── queries/          # SQL query files for SQLC
│   └── tests/            # Test data SQL files
├── docs/                 # Documentation
│   ├── database.md       # Database schema documentation
│   └── environment_variables.md # Environment variables documentation
├── internal/             # Internal packages
│   ├── api/              # API and business logic
│   ├── auth/             # Authentication logic
│   └── database/         # Generated SQLC code
├── scripts/              # Helper scripts
│   └── manage_test_data.sh
├── .env                  # Environment variables (git-ignored)
├── main.go               # Application entry point
├── sqlc.yaml             # SQLC configuration
├── go.mod                # Go module definition
└── go.sum                # Go module checksums
```

---
## 🏛️ License

TODO: Add license information.
