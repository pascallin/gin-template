package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-server/internal/app/todo"
	databases "github.com/pascallin/gin-server/internal/pkg/db"
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
