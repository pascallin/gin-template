package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/pascallin/go-web/docs"
	"github.com/pascallin/go-web/internal/app/task"
	"github.com/pascallin/go-web/internal/app/todo"
	"github.com/pascallin/go-web/internal/app/user"
	databases "github.com/pascallin/go-web/internal/pkg/db"
)

var err error

// @title Gin API
// @version 1.0
// @description A Gin server demo API


// @contact.name pascal_lin

// @host localhost
// @BasePath /v1
func main() {
	// load .env
	godotenv.Load()

	// init db
	// connect mysql
	databases.InitMysqlDatabase()
	defer databases.MysqlDB.Close()
	// connect mongodb
	mongo, err := databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	// initServer
	r := gin.Default()

	v1 := r.Group("/v1")
	task.RegisterRoutes(v1)
	todo.RegisterRoutes(v1)
	user.RegisterRoutes(v1)

	// init swagger
	url := ginSwagger.URL("http://" + os.Getenv("URL") + ":" + os.Getenv("PORT") +"/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// running
	r.Run()
}
