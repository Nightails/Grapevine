package main

import (
	"budgetcli/app"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	p := tea.NewProgram(app.InitialLoginScreenModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
