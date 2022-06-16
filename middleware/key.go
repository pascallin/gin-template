package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pascallin/gin-template/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqKey := c.Request.Header.Get("X-Api-Key")

		if config.GetAppConfig().AppApiKey != reqKey {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
