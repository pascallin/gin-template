package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	Config "github.com/pascallin/go-web/config"
	Models "github.com/pascallin/go-web/models"
	Routes "github.com/pascallin/go-web/routes"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()

	// running
	r.Run()
}
