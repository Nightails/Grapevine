package main

import (
	"budgetcli/internal/cli"
	"budgetcli/internal/config"
	"database/sql"
	"log"

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
	// Run CLI
	cli.Run(cfg, db)
}
