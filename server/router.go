package server

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/pascallin/gin-template/controller"
	"github.com/pascallin/gin-template/middleware"
	"github.com/pascallin/gin-template/ws"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	health := new(HealthController)

	router.GET("/health", health.Status)
	router.GET("metrics", prometheusHandler())

	// websocket serve
	router.GET("/ws", gin.Logger(), func(ctx *gin.Context) {
		ws.ServeWs(ctx)
	})

	// html rendering
	publicPath := filepath.Join(".", "public", "*.tmpl")
	router.LoadHTMLGlob(publicPath)

	v1 := router.Group("/api/v1", gin.Logger(), middleware.I18n(), middleware.ErrorHandler())
	{
		userGroup := v1.Group("user")
		{
			user := new(controller.AuthController)
			userGroup.POST("/login", user.LoginRoute)
			userGroup.POST("/register", user.RegisterRoute)
			userGroup.PATCH("/password", middleware.AuthMiddleware(), user.PatchPasswordRoute)
		}
		fileGroup := v1.Group("files")
		fileGroup.Use(middleware.AuthMiddleware())
		{
			file := new(controller.FileController)
			fileGroup.POST("/upload/xlsx", file.UploadFile)
		}
		taskGroup := v1.Group("task")
		taskGroup.Use(middleware.AuthMiddleware())
		{
			task := new(controller.TaskController)
			taskGroup.GET("/", task.GetTasks)
			taskGroup.POST("/", task.CreateTask)
			taskGroup.GET("/:id", task.GetTask)
			taskGroup.PUT("/:id", task.UpdateTask)
			taskGroup.DELETE("/:id", task.DeleteTask)
		}
		todoGroup := v1.Group("todo")
		todoGroup.Use(middleware.AuthMiddleware())
		{
			todo := new(controller.TodoController)
			todoGroup.GET("/", todo.GetTodos)
			todoGroup.POST("/", todo.CreateTodo)
			todoGroup.GET("/:id", todo.GetTodo)
			todoGroup.PUT("/:id", todo.UpdateTodo)
			todoGroup.DELETE("/:id", todo.DeleteTodo)
		}
		mqGroup := v1.Group("mq")
		{
			mqGroup.POST("/", controller.SendHelloRoute)
		}
	}
	return router

}
