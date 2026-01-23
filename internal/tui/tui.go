package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m GrapeModel) Init() tea.Cmd {
	return nil
}

func (m GrapeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m GrapeModel) View() string {
	return "Welcome to grapevine!"
}
