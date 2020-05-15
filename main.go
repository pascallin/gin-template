package main

import (
	"github.com/joho/godotenv"
	Databases "github.com/pascallin/go-web/databases"
	Routes "github.com/pascallin/go-web/routes"
)

var err error

func main() {

	// load .env
	godotenv.Load()

	// connect mysql
	Databases.InitMysqlDatabase()
	defer Databases.MysqlDB.Close()
	// connect mongodb
	mongo, err := Databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	r := Routes.SetupRouter()

	// running
	r.Run()
}
