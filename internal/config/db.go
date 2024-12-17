package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grabit-api-go/internal/models"
	"log"
	"os"
	"fmt"
)

var db *gorm.DB

// Initialise the database connection
func InitDB() {
	var err error

	user := os.Getenv("DEV_DB_USER")
	password := os.Getenv("DEV_DB_PASSWORD")
	host := os.Getenv("DEV_DB_HOST")
	port := os.Getenv("DEV_DB_PORT")
	dbname := os.Getenv("DEV_DB_NAME")
	sslmode := os.Getenv("DEV_DB_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
	user, password, host, port, dbname, sslmode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{}, 
		&models.UserProfile{},
		&models.Category{}, 
		&models.Ad{}, 
		&models.AdImage{}, 
		&models.Favorite{}, 
		&models.Tag{},
		&models.Message{}, 
		&models.UserActivity{},
	)
}