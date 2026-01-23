package tui

import (
	"database/sql"
	"grapevine/internal/config"
	"grapevine/internal/database"
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
}

func NewGrapeModel(cfg *config.Config, db *sql.DB) GrapeModel {
	return GrapeModel{
		ViewStat:  UserView,
		Config:    cfg,
		DbQueries: database.New(db),
	}
}
