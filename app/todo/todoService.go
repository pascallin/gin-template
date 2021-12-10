package todo

import (
	"time"

	"github.com/pascallin/gin-template/pkg"
)

func getAllTodo(todo *[]Todo, page uint64, pageSize uint64) (err error) {
	if err = pkg.MysqlDB.Order("updated_at desc").Offset(pageSize * (page - 1)).Limit(pageSize).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func getTodo(todo *Todo, id string) (err error) {
	if err := pkg.MysqlDB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func createTodo(todo *Todo) (err error) {
	if err = pkg.MysqlDB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func updateTodo(todo *Todo) (err error, rows int64) {
	result := pkg.MysqlDB.Model(&todo).Updates(Todo{Title: todo.Title, Description: todo.Description, GormModel: pkg.GormModel{UpdatedAt: time.Now()}})
	return result.Error, result.RowsAffected
}

func deleteTodo(todo *Todo, id uint64) (err error, rows int64) {
	result := pkg.MysqlDB.Where("id = ?", id).Delete(todo)
	return result.Error, result.RowsAffected
}
