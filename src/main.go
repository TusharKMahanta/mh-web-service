package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/", home)
	server.Run(":8080")
}

func home(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Welcome to the home page!",
		"body": map[string]interface{}{
			"content": "This is the body content of the home page.",
			"status":  "active",
			"tags":    []string{"home", "welcome", "gin"},
		},
	})
}
