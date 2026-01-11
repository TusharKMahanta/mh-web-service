package router

import (
	"morhaat.com/mh-ui-service/src/config"
	"morhaat.com/mh-ui-service/src/handler"
	"morhaat.com/mh-ui-service/src/middleware"
	"morhaat.com/mh-ui-service/src/utils"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	cfg := config.GetConfig()

	server := gin.Default()
	server.GET("/", middleware.AuthMiddleware(), handler.HomeHandler)

	utils.Log.Info("Starting server on port " + cfg.Server.Port + " in " + cfg.Server.Env + " mode")
	server.Run(":" + cfg.Server.Port)
	return server
}
