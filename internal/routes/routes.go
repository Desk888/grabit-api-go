package routes

import (
	"github.com/gin-gonic/gin"
	// "grabit-api-go/config"
	// "grabit-api-go/models"
	"grabit-api-go/internal/handlers"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/init", handlers.GetInit)
}

