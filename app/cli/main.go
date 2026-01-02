package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Welcome to Budget CLI")

	// load .env file
	fmt.Println("Loading .env file...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("GOOSE_DRIVER")
	dbSource := os.Getenv("GOOSE_DBSTRING")
	if dbSource == "" {
		log.Fatal("GOOSE_DBSTRING environment variable is not set in .env")
	}

	// open database connection
	fmt.Printf("Opening database connection: %s\n", dbSource)
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Verify the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}
