package tui

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"grapevine/internal/auth"
	"grapevine/internal/config"
	"grapevine/internal/database"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

// LoginModel entrypoint to the application. Login or register a user.
type LoginModel struct {
	dbQueries *database.Queries
	overview  OverviewModel

	// Overflow by 1 compare to the length of textInputs, to include login and register buttons
	focusIndex int
	textInputs []textinput.Model
	errorMsg   string
}

func NewLoginModel(dbQueries *database.Queries) LoginModel {
	cfg := config.Config{}
	m := LoginModel{
		dbQueries:  dbQueries,
		overview:   NewOverviewModel(&cfg, dbQueries),
		focusIndex: 0,
		textInputs: make([]textinput.Model, 2),
		errorMsg:   "",
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

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		// Submit form or next focus
		case "enter":
			// Submit
			if m.focusIndex == len(m.textInputs) {
				u, err := m.dbQueries.GetUserByUsername(context.Background(), m.textInputs[0].Value())
				if err != nil {
					m.errorMsg = "User not found"
					cmd := m.updateInputs(msg)
					return m, cmd
				}
				if !auth.CheckPassword(m.textInputs[1].Value(), u.Password) {
					m.errorMsg = "Invalid password"
					cmd := m.updateInputs(msg)
					return m, cmd
				}
				m.overview.cfg.User = u
				return m.overview, nil
			}
			// Register
			if m.focusIndex == len(m.textInputs)+1 {
				hashedPassword, err := auth.HashPassword(m.textInputs[1].Value())
				if err != nil {
					return m, nil
				}
				params := database.AddUserParams{
					ID:        uuid.NewString(),
					CreatedAt: time.Now().String(),
					UpdatedAt: time.Now().String(),
					Username:  m.textInputs[0].Value(),
					Password:  hashedPassword,
				}
				u, err := m.dbQueries.AddUser(context.Background(), params)
				if err != nil {
					log.Printf("error: %v", err)
				} else {
					log.Printf("user: %+v", u)
				}
				m.overview.cfg.User = u
				return m.overview, nil
			}
			m.focusIndex++
			cmds := m.updateFocus()
			return m, tea.Batch(cmds...)
		// Cycle focus
		case "tab", "shift+tab", "up", "down":
			s := msg.String()

			if s == "tab" || s == "down" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}

			if m.focusIndex > len(m.textInputs)+1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.textInputs) + 1
			}

			cmds := m.updateFocus()

			return m, tea.Batch(cmds...)
		}
	case error:
		return m, nil
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m LoginModel) View() string {
	var b strings.Builder

	b.WriteString("Welcome to Grapevine!\n\n")

	if m.errorMsg != "" {
		_, _ = fmt.Fprintf(&b, "%s\n\n", errorStyle.Render(m.errorMsg))
	}

	for i := range m.textInputs {
		b.WriteString(m.textInputs[i].View())
		if i != len(m.textInputs)-1 {
			b.WriteString("\n")
		}
	}

	loginButton := fmt.Sprintf("[ %s ]", blurredStyle.Render("Login"))
	if m.focusIndex == len(m.textInputs) {
		loginButton = fmt.Sprintf("[ %s ]", focusedStyle.Render("Login"))
	}

	registerButton := fmt.Sprintf("[ %s ]", blurredStyle.Render("Register"))
	if m.focusIndex == len(m.textInputs)+1 {
		registerButton = fmt.Sprintf("[ %s ]", focusedStyle.Render("Register"))
	}
	_, _ = fmt.Fprintf(&b, "\n\n%s  %s\n\n", loginButton, registerButton)

	return b.String()
}

func (m LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.textInputs))

	for i := range m.textInputs {
		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m LoginModel) updateFocus() []tea.Cmd {
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

	return cmds
}
