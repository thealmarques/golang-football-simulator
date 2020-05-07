package mock

import (
	"MCE/models"
	"os"
)

// Postgres configuration
var (
	host     = os.Getenv("DB_URL")
	port     = os.Getenv("DB_PORT")
	user     = "admin"
	password = "admin"
	dbname   = "postgres"
)

// Test data interfaces to retrieve the data
type Players struct {
	Players []models.Player `json:"players"`
}

type Teams struct {
	Teams []models.Teams `json:"teams"`
}
