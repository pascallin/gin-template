package main

import (
	"github.com/joho/godotenv"
	internal "github.com/pascallin/go-web/internal"
)

var err error

func main() {
	// load .env
	godotenv.Load()
	// init db
	internal.InitDB()
	r := initServer()
	// running
	r.Run()
}
