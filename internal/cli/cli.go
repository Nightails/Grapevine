package cli

import (
	"budgetcli/internal/config"
	"budgetcli/internal/database"
	"database/sql"
	"fmt"
	"os"
)

type state struct {
	dbQueries *database.Queries
	cfg       *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func Run(cfg *config.Config, db *sql.DB) {
	//s := state{dbQueries: database.New(db), cfg: cfg}
	//cmds := commands{cmdMap: make(map[string]func(*state, command) error)}

	// Parse arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: missing command")
		printHelp()
		os.Exit(1)
	}
}
