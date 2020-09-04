package task

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getTasks(ctx *gin.Context) {
	tasks := GetTasksData()
	if tasks == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, tasks)
	}
}

func createTask(ctx *gin.Context) {
	var task = Task{}
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("create task error"))
	}
	result := CreateTaskData(task.New())
	if result == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("create task error"))
	} else {
		ctx.JSON(http.StatusOK, task)
	}
}

func getTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	result := GetTaskById(id)
	if result == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func updateTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	var task Task
	ctx.BindJSON(&task)
	result := UpdateTaskData(id, &task)
	if result == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func removeTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	err = RemoveTaskData(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "okay"})
	}
}
