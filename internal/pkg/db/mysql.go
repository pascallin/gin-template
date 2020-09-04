package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

const (
	mysqlConnStringTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

type GormModel struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

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
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	MysqlDB = db
}
