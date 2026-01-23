package cli

import (
	"budgetcli/internal/config"
	"budgetcli/internal/database"
	"database/sql"
	"errors"
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
	s := state{dbQueries: database.New(db), cfg: cfg}
	cmds := commands{cmdMap: make(map[string]func(*state, command) error)}

	// Parse arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: missing command")
		printHelp()
		os.Exit(1)
	}

	// Run command
	cmd := command{name: args[1], args: args[2:]}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}

func (cmds *commands) run(s *state, c command) error {
	fn, ok := cmds.cmdMap[c.name]
	if !ok {
		return errors.New("unknown command")
	}
	return fn(s, c)
}
