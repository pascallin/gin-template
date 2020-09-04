package main

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/go-web/internal"
)

func main() {
	// load .env
	godotenv.Load()
	internal.MigrateDB()
}