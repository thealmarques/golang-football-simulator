package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func SetupModels() *gorm.DB {

	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=admin dbname=postgres password=admin sslmode=disable", os.Getenv("DB_URL"), os.Getenv("DB_PORT")))

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	db.AutoMigrate(&Simulation{})

	return db
}
