package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-template/app/todo"
	"github.com/pascallin/gin-template/pkg"
)

func init() {
	godotenv.Load()
}

func main() {
	// connect mysql
	defer pkg.MysqlDB.Close()
	// migration
	pkg.MysqlDB.AutoMigrate(&todo.Todo{})
}
