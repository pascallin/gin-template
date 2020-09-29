package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/go-web/internal/app/todo"
	databases "github.com/pascallin/go-web/internal/pkg/db"
)

func main() {
	// load .env
	godotenv.Load()
	// connect mysql
	databases.InitMysqlDatabase()
	defer databases.MysqlDB.Close()
	// migration
	databases.MysqlDB.AutoMigrate(&todo.Todo{})
}