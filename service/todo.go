package service

import (
	"time"

	"github.com/pascallin/gin-template/conn"
	"github.com/pascallin/gin-template/model"
)

var getMysqlDb = conn.GetMysqlDB

func GetAllTodo(todo *[]model.Todo, page int, pageSize int) (err error) {
	if err = getMysqlDb().Order("updated_at desc").Offset(pageSize * (page - 1)).Limit(pageSize).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *model.Todo, id string) (err error) {
	if err := getMysqlDb().Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(title string, description string) (err error) {
	todo := model.Todo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err = getMysqlDb().Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *model.Todo) (rows int64, err error) {
	result := getMysqlDb().Model(&todo).Updates(model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		UpdatedAt:   time.Now(),
	})
	return result.RowsAffected, result.Error
}

func DeleteTodo(todo *model.Todo, id uint64) (rows int64, err error) {
	result := getMysqlDb().Where("id = ?", id).Delete(todo)
	return result.RowsAffected, result.Error
}
