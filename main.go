package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/go-web/databases"
	"github.com/pascallin/go-web/routes"
)

var err error

func main() {

	// load .env
	godotenv.Load()

	// connect mysql
	databases.InitMysqlDatabase()
	defer databases.MysqlDB.Close()
	// connect mongodb
	mongo, err := databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	r := Routes.SetupRouter()

	// running
	r.Run()
}
