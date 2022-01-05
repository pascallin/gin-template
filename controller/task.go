package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/conn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskController struct{}

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

// @Summary get tasks
// @Description get tasks
// @Tags task
// @Accept  json
// @Produce json
// @Success 200 {array} Task
// @Router /task/ [get]
func (t TaskController) GetTasks(c *gin.Context) {
	input := GetTaskListInput{}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, tasks := getTasksData(findTasksCond{Title: input.Title}, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (t TaskController) GetTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := getTaskById(id)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (t TaskController) CreateTask(c *gin.Context) {
	var task = CreateTaskInput{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, id := createTaskData(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (t TaskController) UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task UpdateTaskInput
	c.BindJSON(&task)
	err, result := updateTaskData(id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no task was found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (t TaskController) DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = removeTaskData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "okay"})
}

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

	cur, err := conn.MongoDB.DB.Collection("tasks").Find(ctx, condition, findOptions)
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
	err := conn.MongoDB.DB.Collection("tasks").FindOne(context.Background(), condition).Decode(&task)
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
	insertResult, err := conn.MongoDB.DB.Collection("tasks").InsertOne(ctx, task)
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
	err := conn.MongoDB.DB.Collection("tasks").FindOneAndUpdate(ctx, filter, update, &opt).Decode(&updatedTask)
	if err != nil {
		return err, nil
	}
	return nil, &updatedTask
}

func removeTaskData(id primitive.ObjectID) error {
	_, err := conn.MongoDB.DB.Collection("tasks").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}
