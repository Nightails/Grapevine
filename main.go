package main

import (
	"database/sql"
	"grapevine/internal/config"
	"grapevine/internal/tui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize config and database
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	db, err := sql.Open(cfg.DbDriver, cfg.DbURl)
	if err != nil {
		log.Fatalf("Failed to open SQL database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close SQL database: %v", err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQL database: %v", err)
	}

	// Initialize TUI
	m := tui.NewGrapeModel(cfg, db)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Failed to run TUI: %v", err)
	}
}
