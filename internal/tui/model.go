package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle      = lipgloss.NewStyle()
)

type Model struct {
	focusIndex int
	textInputs []textinput.Model
}

func InitModel() Model {
	m := Model{
		focusIndex: 0,
		textInputs: make([]textinput.Model, 2),
	}

	for i := range m.textInputs {
		t := textinput.New()
		t.CharLimit = 32
		t.Width = 32

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '*'
		}

		m.textInputs[i] = t
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.textInputs) {
				// TODO: look up username and password in database
				return m, tea.Quit
			}

			if s == "tab" || s == "down" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}

			if m.focusIndex > len(m.textInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.textInputs)
			}

			cmds := make([]tea.Cmd, len(m.textInputs))
			for i := 0; i <= len(m.textInputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.textInputs[i].Focus()
					m.textInputs[i].PromptStyle = focusedStyle
					m.textInputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.textInputs[i].Blur()
				m.textInputs[i].PromptStyle = noStyle
				m.textInputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	case error:
		return m, nil
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.textInputs))

	for i := range m.textInputs {
		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Welcome to Grapevine!\n\n")

	for i := range m.textInputs {
		b.WriteString(m.textInputs[i].View())
		if i != len(m.textInputs)-1 {
			b.WriteString("\n")
		}
	}

	button := fmt.Sprintf("[ %s ]", blurredStyle.Render("Login"))
	if m.focusIndex == len(m.textInputs) {
		button = fmt.Sprintf("[ %s ]", focusedStyle.Render("Login"))
	}
	_, _ = fmt.Fprintf(&b, "\n\n%s\n\n", button)

	return b.String()
}
