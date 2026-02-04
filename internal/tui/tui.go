package tui

import (
	"grapevine/internal/database"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	dbQueries    *database.Queries
	currentState string
	loginModel   loginModel
}

func NewModel(dbQueries *database.Queries) Model {
	return Model{
		dbQueries:    dbQueries,
		currentState: "login",
		loginModel:   newLoginModel(dbQueries),
	}
}

func (m Model) Init() tea.Cmd {
	return m.loginModel.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	lm, cmd := m.loginModel.Update(msg)
	return lm, cmd
}

func (m Model) View() string {
	return m.loginModel.View()
}
