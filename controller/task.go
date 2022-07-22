package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct{}

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
// @Security ApiKeyAuth
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
	tasks, err := service.GetTasksData(service.FindTasksCond{Title: input.Title}, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// @Summary get task
// @Description get task
// @Tags task
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} Task
// @Router /task/:id [get]
// @Param   id     path    string     true        "ID"
func (t TaskController) GetTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := service.GetTaskById(id)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// CreateTask godoc
// @Summary create a task item
// @Schemes
// @Description create a task item
// @Tags task
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Router /task [post]
// @Param   data     body    CreateTaskInput     true        "data"
func (t TaskController) CreateTask(c *gin.Context) {
	var task = CreateTaskInput{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.CreateTaskData(task.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// UpdateTask godoc
// @Summary update a task item
// @Schemes
// @Description update a task item
// @Tags task
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Router /task/:id [put]
// @Param   id     path    string     true        "ID"
// @Param   data     body    UpdateTaskInput     true        "data"
func (t TaskController) UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task UpdateTaskInput
	c.BindJSON(&task)
	result, err := service.UpdateTaskData(id, task.Title)
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

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description delete a task
// @Tags task
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Router /task/:id [delete]
// @Param   id     path    string     true        "ID"
func (t TaskController) DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.RemoveTaskData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "okay"})
}
