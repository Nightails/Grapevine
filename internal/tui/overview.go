package tui

import (
	"fmt"
	"grapevine/internal/config"
	"grapevine/internal/database"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type OverviewModel struct {
	dbQueries *database.Queries
	cfg       *config.Config

	transaction TransactionModel
}

func NewOverviewModel(cfg *config.Config, dbQueries *database.Queries) OverviewModel {
	return OverviewModel{
		cfg:       cfg,
		dbQueries: dbQueries,

		transaction: NewTransactionModel(cfg, dbQueries),
	}
}

func (m OverviewModel) Init() tea.Cmd {
	return nil
}

func (m OverviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "a":
			return m.transaction, nil
		}
	}
	return m, nil
}

func (m OverviewModel) View() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Overview\nWelcome %s!\n\n", m.cfg.User.Username))
	b.WriteString(fmt.Sprintf("Transactions\n"))
	b.WriteString(fmt.Sprintf("Account Balances\n"))
	return b.String()
}
