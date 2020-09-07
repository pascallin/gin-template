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


