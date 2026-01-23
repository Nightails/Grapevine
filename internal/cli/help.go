package cli

import "fmt"

func printHelp() {
	fmt.Println("Usage: grapevine <command> [args...]")
	fmt.Println("Available commands:")
	fmt.Println("  help - Display this help message")
	fmt.Println("  add - Add a new transaction")
	fmt.Println("  list - List all transactions")
	fmt.Println("  delete - Delete a transaction")
}
