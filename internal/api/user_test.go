package api

import (
	"budgetcli/internal/database"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) (*Config, *sql.DB) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	// Create users table for testing
	_, err = db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		t.Fatal(err)
	}

	queries := database.New(db)
	return &Config{DbQueries: queries}, db
}

func TestAddNewUser(t *testing.T) {
	cfg, db := setupTestDB(t)
	defer db.Close()

	t.Run("Successfully add user", func(t *testing.T) {
		username := "testuser"
		password := "securepassword"

		id, err := cfg.AddNewUser(username, password)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if id == "" {
			t.Error("expected a non-empty user ID")
		}

		// Verify user exists in DB
		user, err := cfg.DbQueries.GetUserByUsername(t.Context(), username)
		if err != nil {
			t.Fatalf("failed to fetch user: %v", err)
		}
		if user.Username != username {
			t.Errorf("expected username %s, got %s", username, user.Username)
		}
	})

	t.Run("Fail on duplicate username", func(t *testing.T) {
		username := "duplicate"
		password := "pass"

		_, _ = cfg.AddNewUser(username, password)
		_, err := cfg.AddNewUser(username, password)

		if err == nil {
			t.Error("expected error when adding duplicate user, got nil")
		}
	})
}

func TestLoginUser(t *testing.T) {
	cfg, db := setupTestDB(t)
	defer db.Close()

	username := "loginuser"
	password := "correct_pass"
	userID, _ := cfg.AddNewUser(username, password)

	t.Run("Successful login", func(t *testing.T) {
		id, err := cfg.LoginUser(username, password)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if id != userID {
			t.Errorf("expected ID %s, got %s", userID, id)
		}
	})

	t.Run("Invalid password", func(t *testing.T) {
		_, err := cfg.LoginUser(username, "wrong_pass")
		if err == nil {
			t.Error("expected error for invalid password, got nil")
		}
	})

	t.Run("Non-existent user", func(t *testing.T) {
		_, err := cfg.LoginUser("ghost", password)
		if err == nil {
			t.Error("expected error for non-existent user, got nil")
		}
	})
}
