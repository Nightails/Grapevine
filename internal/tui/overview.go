package tui

import (
	"grapevine/internal/database"

	tea "github.com/charmbracelet/bubbletea"
)

type OverviewModel struct {
	dbQueries *database.Queries
}

func NewOverviewModel(dbQueries *database.Queries) OverviewModel {
	return OverviewModel{dbQueries: dbQueries}
}

func (m OverviewModel) Init() tea.Cmd {
	return nil
}

func (m OverviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m OverviewModel) View() string {
	return "Overview"
}
