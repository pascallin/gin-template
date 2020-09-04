package task

import (
	"context"
	"fmt"

	databases "github.com/pascallin/go-web/internal/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetTasksData() []*Task {
	var results []*Task
	ctx := context.Background()

	condition := bson.D{}

	findOptions := options.Find()
	findOptions.SetLimit(2)
	//findOptions.SetSkip()
	//findOptions.SetSort()

	cur, err := databases.MongoDB.DB.Collection("tasks").Find(ctx, condition, findOptions)
	if err != nil {
		return nil
	}
	fmt.Printf("cur: %+v\n", cur)
	// Close the cursor once finished
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return nil
		}
		results = append(results, &task)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func GetTaskById(id primitive.ObjectID) *Task {
	var task Task
	condition := bson.D{{"_id", id}}
	err := databases.MongoDB.DB.Collection("tasks").FindOne(context.Background(), condition).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

func CreateTaskData(task *Task) *Task {
	ctx := context.Background()
	insertResult, err := databases.MongoDB.DB.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		return nil
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return task
}

func UpdateTaskData(id primitive.ObjectID, task *Task) *Task {
	ctx := context.Background()
	filter := bson.D{{"_id", id}}
	update := bson.M{
		"$set": bson.M{"title": task.Title},
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	fmt.Printf("%v\n", filter)
	fmt.Printf("%v\n", update)

	var updatedTask Task
	err := databases.MongoDB.DB.Collection("tasks").FindOneAndUpdate(ctx, filter, update, &opt).Decode(&updatedTask)
	if err != nil {
		return nil
	}
	return &updatedTask
}

func RemoveTaskData(id primitive.ObjectID) error {
	_, err := databases.MongoDB.DB.Collection("tasks").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}
