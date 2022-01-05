package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqKey := c.Request.Header.Get("X-Api-Key")

		// TODO: get key from env
		var key string

		if key != reqKey {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
