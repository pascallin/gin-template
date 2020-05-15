package Controllers
import (
	"errors"
	"github.com/gin-gonic/gin"
	Models "github.com/pascallin/go-web/models"
	"github.com/pascallin/go-web/repositries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetTasks(ctx *gin.Context) {
	tasks := repositries.GetTasks()
	if tasks == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, tasks)
	}
}

func CreateTask(ctx *gin.Context) {
	var task = Models.Task{}
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("create task error"))
	}
	result := repositries.CreateTask(task.New())
	if result == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("create task error"))
	} else {
		ctx.JSON(http.StatusOK, task)
	}
}

func GetTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	result := repositries.GetTaskById(id)
	if result == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	var task Models.Task
	ctx.BindJSON(&task)
	result := repositries.UpdateTask(id, &task)
	if result == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func RemoveTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}
	err = repositries.RemoveTask(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, bson.D{{"id", id}})
	}
}