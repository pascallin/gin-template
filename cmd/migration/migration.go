package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-template/conn"
	"github.com/pascallin/gin-template/model"
)

func init() {
	godotenv.Load()
}

func main() {
	// migration
	conn.GetMysqlDB().AutoMigrate(&model.Todo{})
}
