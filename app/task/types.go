package task

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
}

func (t *Task) New() *Task {
	return &Task{
		ID:    primitive.NewObjectID(),
		Title: t.Title,
	}
}

type CreateTaskInput struct {
	Title string `binding:"required"`
}

type UpdateTaskInput struct {
	Title string `form:"title"`
}

type GetTaskListInput struct {
	Pagination
	Title string `form:"title"`
}

type Pagination struct {
	PageSize uint64 `form:"pageSize" binding:"required,max=20"`
	Page     uint64 `form:"page" binding:"required,max=100"`
}
