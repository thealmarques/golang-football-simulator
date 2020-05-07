package controllers

import (
	"MCE/models"
	"MCE/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindPlayer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var player models.Player
	player_id := c.Param("id")
	if err := db.Where("id = ?", player_id).First(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Match not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": player})
}

func CreatePlayer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input schemas.CreatePlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create player
	player := models.Player{
		Name:        input.Name,
		Age:         input.Age,
		GoalKeeping: input.GoalKeeping,
		Attack:      input.Attack,
		Defense:     input.Defense,
		Creativity:  input.Creativity,
	}

	db.Create(&player)

	c.JSON(http.StatusOK, gin.H{"data": player})
}
