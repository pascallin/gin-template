package main

import (
"github.com/gin-gonic/gin"
internal "github.com/pascallin/go-web/internal"
)

func initServer() *gin.Engine {
	r := gin.Default()
	// r := gin.New()

	// middleware
	r.Use(logger())

	v1 := r.Group("/v1")
	internal.RegisterRoutes(v1)
	return r
}
