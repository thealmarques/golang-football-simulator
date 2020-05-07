package interfaces

import "MCE/models"

type Team struct {
	ID      uint
	Name    string
	Players []models.Player
}
