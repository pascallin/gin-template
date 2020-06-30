package repositries

import (
	"fmt"

	"github.com/pascallin/go-web/databases"
	"github.com/pascallin/go-web/models"
)

func GetAllTodo(todo *[]Models.Todo) (err error) {
	if err = databases.MysqlDB.Order("created desc").Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Models.Todo) (err error) {
	if err = databases.MysqlDB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Models.Todo, id string) (err error) {
	if err := databases.MysqlDB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Models.Todo, id string) (err error) {
	fmt.Println(todo)
	databases.MysqlDB.Save(todo)
	return nil
}

func DeleteATodo(todo *Models.Todo, id string) (err error) {
	databases.MysqlDB.Where("id = ?", id).Delete(todo)
	return nil
}
