package main

import (
	"database/sql"
	"grapevine/internal/database"
	"grapevine/internal/tui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver = "sqlite3"
	dbSource = "database/app.db"
)

func main() {
	dbQueries := initDatabase()

	// Initialize TUI
	m := tui.NewLoginModel(dbQueries)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Failed to run TUI: %v", err)
	}
}

func initDatabase() *database.Queries {
	log.Println("Initializing database...")
	db, err := sql.Open(dbDriver, dbSource)
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
	log.Println("Database initialized")

	return database.New(db)
}
