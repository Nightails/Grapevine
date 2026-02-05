# 🍇 Grapevine

![Project Status](https://img.shields.io/badge/status-early%20development-orange)
![Go Coverage](https://img.shields.io/badge/coverage-3.7%25-red)

> [!WARNING]
> This project is in early development and is subject to frequent breaking changes. Use with caution.

A Terminal User Interface (TUI) application for budget management.

Grapevine is a friendly terminal companion for tracking your money—quick to use, easy to read, and built to keep your budget organized.

## 🌱 Motivation

I've been using a budgeting app on my phone for a while. It is great, but I wanted something more private and simple that still had the features I was familiar with. Thus,

**Grapevine was born!**

It is straightforward, running in the terminal, and privately stored in an SQLite database that I can back up anywhere. Moreover, my personal banking information is never shared with third parties.

## 🚀 Quick Start

```bash
# Download the latest binary from:
# https://github.com/Nightails/Grapevine/releases

# Linux/macOS:
chmod +x grapevine
./grapevine
```

## 📖 Usage

Grapevine is operated through its Terminal UI (TUI) and is designed around local-first budgeting.

### 🌿 First run

- On first launch, Grapevine creates its configuration directory and file at:
  - `~/.config/grapevine/bgcliconfig.yaml`

### ⌨️ Terminal UI basics

#### 🧭 Common key bindings

- `Tab` / `Shift+Tab`: move focus forward/backward between fields and actions
- `Up` / `Down`: move focus between fields and actions
- `Enter`: activate the focused action
- `Ctrl+C`: quit

> Note: Some screens and workflows may change while the project is in early development.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### 🛠️ Development

#### 🧰 Run from source

**Dependencies:**
- **[Go](https://go.dev/)**: v1.25+
- **[SQLite](https://www.sqlite.org/)**: for local database storage
- **[Goose](https://github.com/pressly/goose)**: database migrations
- **[SQLC](https://sqlc.dev/)**: generate type-safe Go from SQL

```bash
git clone https://github.com/Nightails/Grapevine.git
cd grapevine
go mod download
go run main.go
```

#### 🧬 Generate SQL code

To generate SQL code after changing queries:
```bash
sqlc generate
```

#### ✅ Tests

To run tests:
```bash
go test ./...
```

### 📜 Scripts

Helper scripts are available in the `scripts/` directory:

- **Manage Test Data**:
  ```bash
  ./scripts/manage_test_data.sh [up|down]
  ```

### 📚 Documentation

- [docs/database.md](docs/database.md)
- [docs/environment_variables.md](docs/environment_variables.md)

---

## 🏛️ License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
