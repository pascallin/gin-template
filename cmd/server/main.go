package main

import (
	"github.com/joho/godotenv"
	databases "github.com/pascallin/go-web/internal/pkg/db"
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
	r := initServer()
	// running
	r.Run()
}
