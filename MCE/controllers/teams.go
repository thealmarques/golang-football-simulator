package controllers

import (
	"MCE/models"
	"MCE/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var team models.Teams
	team_id := c.Param("id")

	if err := db.Where("id = ?", team_id).First(&team).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Match not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

func CreateTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input schemas.CreateTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create team
	team := models.Teams{Players: input.Players, Name: input.Name}
	db.Create(&team)

	c.JSON(http.StatusOK, gin.H{"data": team})
}
