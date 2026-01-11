package handler

import (
	"github.com/gin-gonic/gin"
	"morhaat.com/mh-ui-service/src/config"
	"morhaat.com/mh-ui-service/src/utils"
)

func HomeHandler(context *gin.Context) {
	cfg := config.GetConfig()
	utils.Log.WithFields(map[string]interface{}{
		"endpoint": "/",
		"method":   "GET",
	}).Info("Home endpoint accessed")

	context.JSON(200, gin.H{
		"message": "Welcome to the home page!",
		"body": map[string]interface{}{
			"content": "This is the body content of the home page.",
			"status":  "active",
			"env":     cfg.Server.Env,
			"tags":    []string{"home", "welcome", "gin"},
		},
	})
}
