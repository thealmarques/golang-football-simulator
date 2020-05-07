// models/setup.go

package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	db_url := os.Getenv("DB_URL")
	db_port := os.Getenv("DB_PORT")
	postgres_config := fmt.Sprintf("host=%s port=%s user=admin dbname=postgres password=admin sslmode=disable", db_url, db_port)
	db, err := gorm.Open("postgres", postgres_config)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	db.AutoMigrate(&Matches{}, &Player{}, &Teams{})

	return db
}
