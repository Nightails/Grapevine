package api

import "budgetcli/internal/database"

type Config struct {
	DbQueries *database.Queries
}
