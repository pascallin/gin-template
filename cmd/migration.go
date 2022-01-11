package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-template/conn"
	"github.com/pascallin/gin-template/controller"
)

func init() {
	godotenv.Load()
}

func main() {
	// connect mysql
	defer conn.MysqlDB.Close()
	// migration
	conn.MysqlDB.AutoMigrate(&controller.Todo{})
}
