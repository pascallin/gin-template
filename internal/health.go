package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckRoutes(rg *gin.RouterGroup) {
	rg.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Pong",
		})
	})
}
