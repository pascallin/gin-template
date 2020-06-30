package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pascallin/go-web/api"
	"github.com/pascallin/go-web/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r := gin.New()

	// middleware
	r.Use(Middleware.Logger())

	v1 := r.Group("/v1")
	{
		v1.GET("todo", api.GetTodoList)
		v1.POST("todo", api.CreateATodo)
		v1.GET("todo/:id", api.GetATodo)
		v1.PUT("todo/:id", api.UpdateATodo)
		v1.DELETE("todo/:id", api.DeleteATodo)

		v1.GET("task", api.GetTasks)
		v1.POST("task", api.CreateTask)
		v1.PUT("task/:id", api.UpdateTask)
		v1.DELETE("task/:id", api.RemoveTask)
		v1.GET("task/:id", api.GetTask)
	}
	return r
}
