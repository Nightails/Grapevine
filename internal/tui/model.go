package tui

import (
	"database/sql"
	"grapevine/internal/config"
	"grapevine/internal/database"

	"github.com/charmbracelet/bubbles/textinput"
)

const (
	UserView uint = iota
	LoginView
	SummaryView
)

// GrapeModel The main model for the grapevine TUI
type GrapeModel struct {
	ViewStat  uint
	Config    *config.Config
	DbQueries *database.Queries

	userNameInput textinput.Model
}

func NewGrapeModel(cfg *config.Config, db *sql.DB) GrapeModel {
	ti := textinput.New()
	ti.Placeholder = "Enter your username"
	ti.Focus()
	ti.Width = 20

	return GrapeModel{
		ViewStat:  UserView,
		Config:    cfg,
		DbQueries: database.New(db),

		userNameInput: ti,
	}
}
