package task

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/task")

	router.GET("/", getTasks)
	router.POST("/", createTask)
	router.GET("/:id", getTask)
	router.PUT("/:id", updateTask)
	router.DELETE("/:id", removeTask)
}
