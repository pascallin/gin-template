package task

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getTasks(c *gin.Context) {
	tasks := getTasksData()
	if tasks == nil {
		c.JSON(http.StatusOK, []Task{})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func createTask(c *gin.Context) {
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
	c.JSON(http.StatusOK, bson.M{"id": id})
}

func getTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}
	result := getTaskById(id)
	if result == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func updateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task Task
	c.BindJSON(&task)
	result := updateTaskData(id, &task)
	if result == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func removeTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}
	err = removeTaskData(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "okay"})
	}
}
