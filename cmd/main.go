package main

import (
	"grabit-api-go/internal/config"
	"log"
	"grabit-api-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise Database
	config.InitDB()
	
	// Initialise Gin Router
	r := gin.Default()

	// Register routes
	routes.InitRoutes(r)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}