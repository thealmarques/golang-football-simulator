package controllers

import (
	"MSE/models"
	"MSE/routine"
	"MSE/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateSimulation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input schemas.CreateSimulation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	simulation := models.Simulation{
		MatchId: input.MatchId,
		Result:  "",
		Events:  []string{},
	}

	db.Create(&simulation)

	c.JSON(http.StatusOK, gin.H{"data": simulation})

	// Run the simulation
	go routine.Run(input, simulation, input.MatchId)
}

func UpdateSimulation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// get input json
	var input schemas.UpdateSimulation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get model if exists
	var simulation models.Simulation
	if err := db.Where("id = ?", fmt.Sprint(input.ID)).First(&simulation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Model(&simulation).Update(&input)

	c.JSON(http.StatusOK, gin.H{"data": simulation})
}
