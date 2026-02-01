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
	m := tui.InitModel(dbQueries)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Failed to run TUI program: %v", err)
	}
}

func initDatabase() *database.Queries {
	log.Println("Initializing database...")
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to open SQL database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQL database: %v", err)
	}
	log.Println("Database initialized")

	return database.New(db)
}
