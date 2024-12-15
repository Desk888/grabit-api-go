package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grabit-api-go/models"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error

	dsn := "postgres://postgres:postgres@localhost:5432/grabit?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.UserProfile{}, &models.Ad{}, &models.AdImage{}, &models.Category{}, &models.Favourite{}, &models.Message{}, &models.UserActivity{})
}