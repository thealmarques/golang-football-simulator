package interfaces

import "MSS/models"

type Team struct {
	ID      uint
	Name    string
	Players []models.Player
}
