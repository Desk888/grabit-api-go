package routes

import (
	"fmt"
	// "net/http"
	"github.com/gin-gonic/gin"
	// "grabit-api-go/config"
	// "grabit-api-go/models"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/init", GetInit)
}

func GetInit(c *gin.Context) {
	fmt.Println("Init Route Called")
	c.JSON(200, gin.H{
		"message": "Init Route Called",
	})
}