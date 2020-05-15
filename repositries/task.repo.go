package repositries

import (
	"context"
	"fmt"
	"github.com/pascallin/go-web/databases"
	models "github.com/pascallin/go-web/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTasks() []*models.Task {
	var results []*models.Task
	ctx := context.Background()

	condition :=  bson.D{}

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
		var task models.Task
		err := cur.Decode(&task); if err != nil {
			return nil
		}
		results = append(results, &task)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func GetTaskById(id primitive.ObjectID) *models.Task {
	var task models.Task
	condition := bson.D{{"_id",id}}
	err := databases.MongoDB.DB.Collection("tasks").FindOne(context.Background(), condition).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

func CreateTask(task *models.Task) *models.Task {
	ctx := context.Background()
	insertResult, err := databases.MongoDB.DB.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		return nil
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return task
}

func UpdateTask(id primitive.ObjectID, task *models.Task) *models.Task {
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

	var updatedTask models.Task
	err := databases.MongoDB.DB.Collection("tasks").FindOneAndUpdate(ctx, filter, update, &opt).Decode(&updatedTask)
	if err != nil {
		return nil
	}
	return &updatedTask
}

func RemoveTask(id primitive.ObjectID) error {
	_, err := databases.MongoDB.DB.Collection("tasks").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}