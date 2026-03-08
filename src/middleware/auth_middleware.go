package middleware

import (
	"github.com/gin-gonic/gin"
	"morhaat.com/mh-ui-service/src/utils"
)

var Log = utils.Log

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Log.Info("AuthMiddleware: Checking authentication")
		// Placeholder for actual authentication logic
		// For example, check for a valid token in headers
		c.Next()
	}
}
