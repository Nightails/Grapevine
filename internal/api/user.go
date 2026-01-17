package api

import (
	"budgetcli/internal/auth"
	"budgetcli/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// AddNewUser adds a new user to the database and returns the user ID and any error encountered
func (cfg *Config) AddNewUser(username, password string) (string, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := cfg.DbQueries.AddUser(context.Background(), database.AddUserParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Username:  username,
		Password:  hashedPassword,
	})
	if err != nil {
		return "", fmt.Errorf("failed to add user: %w", err)
	}

	return user.ID, nil
}

// LoginUser authenticates a user and returns their ID and any error encountered
func (cfg *Config) LoginUser(username, password string) (string, error) {
	user, err := cfg.DbQueries.GetUserByUsername(context.Background(), username)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	if !auth.CheckPassword(password, user.Password) {
		return "", fmt.Errorf("invalid password")
	}

	return user.ID, nil
}

func (cfg *Config) IsUserExists(username string) bool {
	_, err := cfg.DbQueries.GetUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	return true
}
