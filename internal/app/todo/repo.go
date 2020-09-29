package todo

import (
	databases "github.com/pascallin/go-web/internal/pkg/db"
	"time"
)

func getAllTodo(todo *[]Todo, page uint64, pageSize uint64) (err error) {
	if err = databases.MysqlDB.Order("updated_at desc").Offset(pageSize * (page - 1)).Limit(pageSize).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func getTodo(todo *Todo, id string) (err error) {
	if err := databases.MysqlDB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func createTodo(todo *Todo) (err error) {
	if err = databases.MysqlDB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func updateTodo(todo *Todo) (err error, rows int64) {
	result := databases.MysqlDB.Model(&todo).Updates(Todo{Title: todo.Title, Description:todo.Description, GormModel: databases.GormModel{UpdatedAt: time.Now()}})
	return result.Error, result.RowsAffected
}

func deleteTodo(todo *Todo, id uint64) (err error, rows int64) {
	result := databases.MysqlDB.Where("id = ?", id).Delete(todo)
	return result.Error, result.RowsAffected
}