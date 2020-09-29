package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pascallin/go-web/internal"
	databases "github.com/pascallin/go-web/internal/db"
)

var err error

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
	internal.RegisterRoutes(v1)

	// running
	r.Run()
}
