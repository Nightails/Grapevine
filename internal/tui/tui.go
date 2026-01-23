package tui

import (
	"database/sql"
	"grapevine/internal/config"
	"grapevine/internal/database"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	UserView uint = iota
	LoginView
	SummaryView
)

type Model struct {
	ViewStat  uint
	Config    *config.Config
	DbQueries *database.Queries
}

func NewModel(cfg *config.Config, db *sql.DB) Model {
	return Model{
		ViewStat:  UserView,
		Config:    cfg,
		DbQueries: database.New(db),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Welcome to grapevine!"
}
