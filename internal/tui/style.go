package tui

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5faf"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#585858"))
	noStyle      = lipgloss.NewStyle()
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#ff0000")).PaddingLeft(1).PaddingRight(1)
)
