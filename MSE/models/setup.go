package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=25432 user=admin dbname=postgres password=admin sslmode=disable")

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	db.AutoMigrate(&Simulation{})

	return db
}
