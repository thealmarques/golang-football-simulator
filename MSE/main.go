package main

import (
	"MSE/controllers"
	"MSE/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := models.SetupModels()
	defer db.Close()

	// middleware - intercept requests to use our db controller
	router.Use(func(context *gin.Context) {
		// provide db variable to controllers
		context.Set("db", db)
		context.Next()
	})

	// simulations
	router.POST("/simulation", controllers.CreateSimulation)
	router.PUT("/simulation", controllers.UpdateSimulation)

	router.Run("localhost:9083")
}
