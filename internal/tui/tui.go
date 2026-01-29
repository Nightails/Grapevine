package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m GrapeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m GrapeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle user input
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// Quit the program on Ctrl+C
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case error:
		return m, nil
	}

	// Handle text input
	m.userNameInput, cmd = m.userNameInput.Update(msg)

	return m, cmd
}

func (m GrapeModel) View() string {
	return fmt.Sprintf("Welcome to grapevine!\n%s", m.userNameInput.View()) + "\n"
}
