package todo

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"

	databases "github.com/pascallin/go-web/internal/pkg/db"
)

type Todo struct {
	gorm.Model
	Title       sql.NullString
	Description string
}

func GetAllTodo(todo *[]Todo) (err error) {
	if err = databases.MysqlDB.Order("created desc").Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err = databases.MysqlDB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	if err := databases.MysqlDB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	databases.MysqlDB.Save(todo)
	return nil
}

func DeleteTodo(todo *Todo, id string) (err error) {
	databases.MysqlDB.Where("id = ?", id).Delete(todo)
	return nil
}
