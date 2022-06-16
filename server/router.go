package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/controller"
	"github.com/pascallin/gin-template/middleware"
	"github.com/pascallin/gin-template/sender"
)

func NewRouter(hub *Hub) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	health := new(HealthController)

	router.GET("/health", health.Status)

	// websocket serve
	router.GET("/ws", gin.Logger(), func(ctx *gin.Context) {
		serveWs(hub, ctx)
	})

	v1 := router.Group("v1", gin.Logger(), middleware.AuthMiddleware())
	{
		userGroup := v1.Group("user")
		{
			user := new(controller.AuthController)
			userGroup.POST("/login", user.LoginRoute)
			userGroup.POST("/register", user.RegisterRoute)
			userGroup.PATCH("/password", user.PatchPasswordRoute)
		}
		fileGroup := v1.Group("files")
		{
			file := new(controller.FileController)
			fileGroup.POST("/upload/xlsx", file.UploadFile)
		}
		taskGroup := v1.Group("task")
		{
			task := new(controller.TaskController)
			taskGroup.GET("/", task.GetTasks)
			taskGroup.GET("/:id", task.GetTask)
			taskGroup.POST("/", task.CreateTask)
			taskGroup.PUT("/:id", task.UpdateTask)
			taskGroup.DELETE("/:id", task.DeleteTask)
		}
		todoGroup := v1.Group("todo")
		{
			todo := new(controller.TodoController)
			todoGroup.GET("/", todo.GetTodos)
			todoGroup.GET("/:id", todo.GetTodo)
			todoGroup.POST("/", todo.CreateTodo)
			todoGroup.PUT("/:id", todo.UpdateTodo)
			todoGroup.DELETE("/:id", todo.DeleteTodo)
		}
		mqGroup := v1.Group("mq")
		{
			mqGroup.POST("/", sender.SendHelloRoute)
		}
		verifyGroup := v1.Group("verify")
		{
			verifyCtrl := new(controller.VerifyController)
			verifyGroup.POST("rtmp", verifyCtrl.AuthOnly)
		}
	}
	return router

}
