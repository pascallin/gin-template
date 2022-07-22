package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pascallin/gin-template/conn"
	"github.com/pascallin/gin-template/model"
)

type FindTasksCond struct {
	Title string `json:"title"`
}

func GetTasksData(cond FindTasksCond, page int, pageSize int) ([]*model.Task, error) {
	var results []*model.Task
	ctx := context.Background()

	condition := bson.D{}
	if cond.Title != "" {
		condition = append(condition, bson.E{Key: "title", Value: primitive.Regex{Pattern: cond.Title, Options: "i"}})
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetSort(bson.M{"title": -1})

	cur, err := conn.GetMongo(ctx).DB.Collection("tasks").Find(ctx, condition, findOptions)
	if err != nil {
		return results, err
	}
	fmt.Printf("cur: %+v\n", cur)
	// Close the cursor once finished
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var task model.Task
		err := cur.Decode(&task)
		if err != nil {
			return results, err
		}
		results = append(results, &task)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil
}

func GetTaskById(id primitive.ObjectID) *model.Task {
	var task model.Task
	condition := bson.M{"_id": id}
	ctx := context.TODO()
	err := conn.GetMongo(ctx).DB.Collection("tasks").FindOne(context.Background(), condition).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

func CreateTaskData(title string) (primitive.ObjectID, error) {
	task := model.Task{
		ID:    primitive.NewObjectID(),
		Title: title,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertResult, err := conn.GetMongo(ctx).DB.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func UpdateTaskData(id primitive.ObjectID, title string) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{
		"$set": bson.M{"title": title},
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	fmt.Printf("%v\n", filter)
	fmt.Printf("%v\n", update)

	var updatedTask model.Task
	err := conn.GetMongo(ctx).DB.Collection("tasks").FindOneAndUpdate(ctx, filter, update, &opt).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func RemoveTaskData(id primitive.ObjectID) error {
	ctx := context.TODO()
	_, err := conn.GetMongo(ctx).DB.Collection("tasks").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}
