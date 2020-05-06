package mock

import "MSS/models"

// Postgres configuration
const (
	host     = "localhost"
	port     = 15432
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
