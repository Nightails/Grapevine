# 🍇 Grapevine

![Project Status](https://img.shields.io/badge/status-early%20development-orange)
![Go Coverage](https://img.shields.io/badge/coverage-3.7%25-red)

> [!WARNING]
> This project is in early development and is subject to frequent breaking changes. Use with caution.

A Terminal User Interface (TUI) application for budget management.

---
## 📝 Overview

Grapevine is a terminal-based tool designed to help users manage their budgets directly from their favorite terminal. It features a rich TUI built with Bubble Tea and Lip Gloss, providing a modern terminal experience. It uses SQLite for local data storage, Goose for database migrations, and SQLC for type-safe database queries.

### ✨ Key Features

- **TUI Dashboard**: A clean and responsive terminal interface.
- **User Authentication**: Simple login system (WIP).
- **Transaction Management**: Track your income and expenses.
- **Account Tracking**: Manage multiple financial accounts.
- **Local First**: All data is stored locally in a SQLite database.

---
## 🛠️ Requirements

- **[Go](https://go.dev/)**: v1.25
- **[SQLite](https://www.sqlite.org/)**: Ensure you have SQLite installed for database storage.
- **[Goose](https://github.com/pressly/goose)**: Used for database migrations.
- **[SQLC](https://sqlc.dev/)**: Used to generate type-safe Go code from SQL.

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

3. **Configure the application**:
   The application automatically creates a configuration directory and file at `~/.config/grapevine/bgcliconfig.yaml` on first run.

   Alternatively, you can refer to [docs/environment_variables.md](docs/environment_variables.md) for details on available environment variables.

4. **Run migrations**:
   ```bash
   # Using goose with explicit flags
   goose -dir database/migrations sqlite3 database/app.db up

   # OR using environment variables (if configured)
   goose up
   ```

5. **Generate SQL code (Optional - for development)**:
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

---
## 📄 Documentation

For detailed information about the database schema and environment variables, please refer to:
- [docs/database.md](docs/database.md)
- [docs/environment_variables.md](docs/environment_variables.md)

---
## 🪧 Tests

Automated Go tests are included in the project, covering authentication and database operations.

To run Go tests:
```bash
go test ./...
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
│   ├── auth/             # Authentication logic
│   ├── config/           # Configuration management
│   ├── database/         # Generated SQLC code and database utilities
│   └── tui/              # Terminal User Interface (bubbletea/lipgloss)
├── scripts/              # Helper scripts
│   └── manage_test_data.sh
├── .env                  # Environment variables (git-ignored)
├── LICENSE               # Project license
├── main.go               # Application entry point
├── sqlc.yaml             # SQLC configuration
├── go.mod                # Go module definition
└── go.sum                # Go module checksums
```

---
## 🏛️ License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
