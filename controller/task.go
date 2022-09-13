package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pascallin/gin-template/model"
	"github.com/pascallin/gin-template/pkg"
	"github.com/pascallin/gin-template/service"
	"github.com/pascallin/gin-template/types"
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

type GetTasksRes struct {
	types.AppResponse
	Data []*model.Task
}
type TaskRes struct {
	types.AppResponse
	Data model.Task
}

// @Summary get tasks
// @Description get tasks
// @Tags task
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} GetTasksRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /task/ [get]
func (t TaskController) GetTasks(c *gin.Context) {
	input := GetTaskListInput{}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.Error(types.ErrParam)
		c.Abort()
		return
	}
	tasks, err := service.GetTasksData(service.FindTasksCond{Title: input.Title}, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, GetTasksRes{
		AppResponse: types.AppResponse{
			Code:    types.SucceedCode,
			Message: pkg.GetI18nMessage(c, types.SucceedCode),
		},
		Data: tasks,
	})
}

// @Summary get task
// @Description get task
// @Tags task
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} TaskRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /task/:id [get]
// @Param   id     path    string     true        "ID"
func (t TaskController) GetTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.Error(types.ErrParam)
		c.Abort()
		return
	}
	result := service.GetTaskById(id)
	if result == nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, pkg.NewSucceedAppResponse(c, result))
}

// CreateTask godoc
// @Summary create a task item
// @Schemes
// @Description create a task item
// @Tags task
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Success 200 {object} TaskRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /task [post]
// @Param   data     body    CreateTaskInput     true        "data"
func (t TaskController) CreateTask(c *gin.Context) {
	var task = CreateTaskInput{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.Error(types.ErrParam)
		c.Abort()
	}
	id, err := service.CreateTaskData(task.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
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
// @Success 200 {object} TaskRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /task/:id [put]
// @Param   id     path    string     true        "ID"
// @Param   data     body    UpdateTaskInput     true        "data"
func (t TaskController) UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.Error(types.ErrParam)
		c.Abort()
	}
	var task UpdateTaskInput
	c.BindJSON(&task)
	result, err := service.UpdateTaskData(id, task.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	if result == nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, "no task was found"))
		return
	}
	c.JSON(http.StatusOK, pkg.NewSucceedAppResponse(c, result))
}

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description delete a task
// @Tags task
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Success 200 {object} types.AppResponse
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /task/:id [delete]
// @Param   id     path    string     true        "ID"
func (t TaskController) DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.Error(types.ErrParam)
		c.Abort()
	}
	err = service.RemoveTaskData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, types.NewAppResponse(types.SucceedCode, "ok"))
}
