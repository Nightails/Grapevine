package main

import (
	"budgetcli/internal/api"
	"budgetcli/internal/database"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver = "sqlite3"
	dbSource = "database/app.db"
)

func main() {
	fmt.Println("Welcome to Budget CLI")

	// open database connection
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

	// Create API config
	config := api.Config{DbQueries: database.New(db)}

	// Prompt for username
	username := getInput("Username")
	if !config.IsUserExists(username) {
		fmt.Println("User does not exist. Please create a new account.")
		os.Exit(1)
	}
	fmt.Println("Welcome, " + username)
}

func getInput(prompt string) string {
	fmt.Print("> " + prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return input
}
