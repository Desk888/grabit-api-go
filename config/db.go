package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grabit-api-go/models"
	"log"
	"os"
	"fmt"
)

var db *gorm.DB

func InitDB() {
	var err error

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
	user, password, host, port, dbname, sslmode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.UserProfile{}, &models.Ad{}, &models.AdImage{}, &models.Category{}, &models.Favourite{}, &models.Message{}, &models.UserActivity{})
}