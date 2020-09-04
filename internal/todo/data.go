package todo

import (
	"fmt"

	databases "github.com/pascallin/go-web/internal/pkg/db"
)

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
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
