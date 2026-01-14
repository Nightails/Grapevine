package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	username = iota
	password
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type LoginScreenModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func InitialLoginScreenModel() LoginScreenModel {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].Prompt = ""
	inputs[username].Width = 30
	inputs[username].Focus()

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "password"
	inputs[password].Prompt = ""
	inputs[password].Width = 30

	return LoginScreenModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m LoginScreenModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case error:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m LoginScreenModel) View() string {
	return fmt.Sprintf(
		`Welcome to Budget CLI!
%s
%s

%s
%s

%s
`,
		inputStyle.Render("Username:"),
		m.inputs[username].View(),
		inputStyle.Render("Password:"),
		m.inputs[password].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput cycles the focused input field to the next one in the inputs slice.
func (m *LoginScreenModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}
