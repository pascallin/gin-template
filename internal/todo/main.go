package todo

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/todo")

	router.GET("/", getTodoList)
	router.POST("/", createATodo)
	router.GET(":id", getATodo)
	router.PUT(":id", updateATodo)
	router.DELETE(":id", deleteATodo)
}

