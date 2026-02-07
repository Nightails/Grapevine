package tui

import (
	"grapevine/internal/config"
	"grapevine/internal/database"

	tea "github.com/charmbracelet/bubbletea"
)

type TransactionModel struct {
	cfg       *config.Config
	dbQueries *database.Queries
}

func NewTransactionModel(cfg *config.Config, dbQueries *database.Queries) TransactionModel {
	return TransactionModel{
		cfg:       cfg,
		dbQueries: dbQueries,
	}
}

func (m TransactionModel) Init() tea.Cmd {
	return nil
}

func (m TransactionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m TransactionModel) View() string {
	return "Transaction"
}
