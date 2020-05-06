// controllers/matches.go

package controllers

import (
	"MSS/models"
	"MSS/requests"
	"MSS/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindMatches(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matches []models.Matches
	db.Order("id DESC").Find(&matches)

	c.JSON(http.StatusOK, gin.H{"data": matches})
}

func CreateMatch(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input schemas.CreateMatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create match
	match := models.Matches{HomeId: input.HomeId, AwayId: input.AwayId}
	db.Create(&match)

	// Get team data
	homeTeam := requests.GetTeam(match.HomeId)
	awayTeam := requests.GetTeam(match.AwayId)

	simulation := requests.PostSimulation(homeTeam, awayTeam, match.ID)

	fmt.Println(simulation)

	c.JSON(http.StatusOK, gin.H{"data": match})
}

func FindMatch(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var match models.Matches
	match_id := c.Param("id")
	if err := db.Where("home_id = ?", match_id).Or("away_id = ?", match_id).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Match not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": match})
}
