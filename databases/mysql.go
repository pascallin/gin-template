package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pascallin/go-web/models"
	"os"
)

const (
	mysqlConnStringTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

func getMysqlConnURL() string {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	dbName := os.Getenv("MYSQL_DATABASE")
	password := os.Getenv("MYSQL_PASSWORD")
	return fmt.Sprintf(
		mysqlConnStringTemplate,
		user,
		password,
		host,
		port,
		dbName,
	)
}

var MysqlDB *gorm.DB

func InitMysqlDatabase() () {
	db, err := gorm.Open("mysql", getMysqlConnURL())
	if err != nil {
		panic(err)
	}
	// TODO: abstract migration
	db.AutoMigrate(&Models.Todo{})
	MysqlDB = db
}
