package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	user "github.com/pascallin/gin-template/app/auth"
	"github.com/pascallin/gin-template/app/task"
	"github.com/pascallin/gin-template/app/todo"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	// load .env
	godotenv.Load()
}

func InitServer() *gin.Engine {
	// initServer
	r := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	task.RegisterRoutes(v1)
	todo.RegisterRoutes(v1)
	user.RegisterRoutes(v1)
	registerHealthCheckRoutes(v1)

	// init swagger
	url := ginSwagger.URL("http://" + os.Getenv("URL") + ":" + os.Getenv("PORT") + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
