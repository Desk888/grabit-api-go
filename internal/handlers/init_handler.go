package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInit(c *gin.Context) {
	fmt.Println("Init Route Called")
	c.JSON(200, gin.H{
		"message": "Init Route Called",
	})
}