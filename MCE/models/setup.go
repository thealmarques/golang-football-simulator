// models/setup.go

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=15432 user=admin dbname=postgres password=admin sslmode=disable")

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	db.AutoMigrate(&Matches{}, &Player{}, &Teams{})

	return db
}
