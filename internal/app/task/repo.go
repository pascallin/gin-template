package task

import (
	"context"
	"fmt"
	"time"

	databases "github.com/pascallin/gin-server/internal/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type findTasksCond struct {
	Title string `json:"title"`
}

func getTasksData(cond findTasksCond, page uint64, pageSize uint64) (error, []*Task) {
	var results []*Task
	ctx := context.Background()

	condition := bson.D{}
	if cond.Title != "" {
		condition = append(condition, bson.E{"title", primitive.Regex{Pattern: cond.Title, Options: "i"}})
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetSort(bson.M{"title": -1})

	cur, err := databases.MongoDB.DB.Collection("tasks").Find(ctx, condition, findOptions)
	if err != nil {
		return err, results
	}
	fmt.Printf("cur: %+v\n", cur)
	// Close the cursor once finished
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return err, results
		}
		results = append(results, &task)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return nil, results
}

func getTaskById(id primitive.ObjectID) *Task {
	var task Task
	condition := bson.M{"_id": id}
	err := databases.MongoDB.DB.Collection("tasks").FindOne(context.Background(), condition).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

func createTaskData(input *CreateTaskInput) (error, primitive.ObjectID) {
	task := Task{
		ID:    primitive.NewObjectID(),
		Title: input.Title,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertResult, err := databases.MongoDB.DB.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		return err, primitive.NilObjectID
	}
	return nil, insertResult.InsertedID.(primitive.ObjectID)
}

func updateTaskData(id primitive.ObjectID, input *UpdateTaskInput) (error, *Task) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	update := bson.M{
		"$set": bson.M{"title": input.Title},
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
		return err, nil
	}
	return nil, &updatedTask
}

func removeTaskData(id primitive.ObjectID) error {
	_, err := databases.MongoDB.DB.Collection("tasks").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}
