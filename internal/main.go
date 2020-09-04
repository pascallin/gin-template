package internal

import (
	"github.com/gin-gonic/gin"
	databases "github.com/pascallin/go-web/internal/pkg/db"
	"github.com/pascallin/go-web/internal/task"
	"github.com/pascallin/go-web/internal/todo"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	task.RegisterRoutes(rg)
	todo.RegisterRoutes(rg)
}

func InitDB() {
	// connect mysql
	databases.InitMysqlDatabase()
	defer databases.MysqlDB.Close()
	// connect mongodb
	mongo, err := databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
}
