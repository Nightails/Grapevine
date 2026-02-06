package tui

import (
	"grapevine/internal/config"
	"grapevine/internal/database"

	tea "github.com/charmbracelet/bubbletea"
)

type OverviewModel struct {
	dbQueries *database.Queries
	cfg       *config.Config
}

func NewOverviewModel(cfg *config.Config, dbQueries *database.Queries) OverviewModel {
	return OverviewModel{
		cfg:       cfg,
		dbQueries: dbQueries,
	}
}

func (m OverviewModel) Init() tea.Cmd {
	return nil
}

func (m OverviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m OverviewModel) View() string {
	return "Overview"
}
