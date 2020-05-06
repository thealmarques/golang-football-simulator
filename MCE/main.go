package main

import (
	"MSS/controllers"
	"MSS/mock"
	"MSS/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mock.CleanDB()

	db := models.SetupModels()
	defer db.Close()

	mock.InsertPlayers()
	mock.InsertTeams()

	// middleware - intercept requests to use our db controller
	router.Use(func(context *gin.Context) {
		// provide db variable to controllers
		context.Set("db", db)
		context.Next()
	})

	// matches
	router.GET("/matches", controllers.FindMatches)
	router.GET("/matches/:id", controllers.FindMatch)
	router.POST("/matches", controllers.CreateMatch)

	// teams
	router.GET("/teams/:id", controllers.FindTeam)
	router.POST("/teams", controllers.CreateTeam)

	// players
	router.GET("/player/:id", controllers.FindPlayer)
	router.POST("/player", controllers.CreatePlayer)

	router.Run("localhost:9082")
}
