package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pascallin/go-web/Controllers"
	Middlewares "github.com/pascallin/go-web/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r := gin.New()

	// middlweares
	r.Use(Middlewares.Logger())

	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)

		v1.GET("task", Controllers.GetTasks)
		v1.POST("task", Controllers.CreateTask)
		v1.PUT("task/:id", Controllers.UpdateTask)
		v1.DELETE("task/:id", Controllers.RemoveTask)
		v1.GET("task/:id", Controllers.GetTask)
	}
	return r
}
