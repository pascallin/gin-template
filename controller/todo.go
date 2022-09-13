package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/model"
	"github.com/pascallin/gin-template/pkg"
	"github.com/pascallin/gin-template/service"
	"github.com/pascallin/gin-template/types"
)

type TodoController struct{}

type Pagination struct {
	PageSize int `form:"pageSize" binding:"required,max=20"`
	Page     int `form:"page" binding:"required,max=100"`
}

type TodoInput struct {
	Title       string `form:"title" xml:"title" json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTodo struct {
	ID          uint64 `uri:"id" binding:"required" json:"id"`
	Title       string `form:"title" xml:"title" json:"title"`
	Description string `form:"description" xml:"description" json:"description"`
}

type GetTodosRes struct {
	types.AppResponse
	Data []model.Todo
}

type TodoRes struct {
	types.AppResponse
	Data model.Todo
}

// GetTodos godoc
// @Summary get todo list
// @Description get todo list
// @Tags todo
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} GetTodosRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /todo [get]
// @Param   page     query    number     true        "page number"
// @Param   pageSize    query    number     true        "page size"
func (t TodoController) GetTodos(c *gin.Context) {
	var input Pagination
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, types.AppResponse{
			Code:    types.ParamErrorCode,
			Message: pkg.GetI18nMessage(c, types.ParamErrorCode),
		})
		return
	}
	var todos []model.Todo
	err := service.GetAllTodo(&todos, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
	} else {
		c.JSON(http.StatusOK, GetTodosRes{
			AppResponse: types.AppResponse{
				Code:    types.SucceedCode,
				Message: pkg.GetI18nMessage(c, types.SucceedCode),
			},
			Data: todos,
		})
	}
}

// GetTodo godoc
// @Summary get todo
// @Description get todo
// @Tags todo
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} TodoRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /todo/:id [get]
// @Param   id     path    string     true        "ID"
func (t TodoController) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, types.AppResponse{
			Code:    types.ParamErrorCode,
			Message: pkg.GetI18nMessage(c, types.ParamErrorCode),
		})
		return
	}
	var todo model.Todo
	err := service.GetTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
	} else {
		c.JSON(http.StatusOK, TodoRes{
			AppResponse: types.AppResponse{
				Code:    types.SucceedCode,
				Message: pkg.GetI18nMessage(c, types.SucceedCode),
			},
			Data: todo,
		})
	}
}

// CreateTodo godoc
// @Summary create a todo item
// @Schemes
// @Description create a todo item
// @Tags todo
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Success 200 {object} TodoRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /todo [post]
// @Param   data     body    TodoInput     true        "data"
func (t TodoController) CreateTodo(c *gin.Context) {
	var input TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, types.AppResponse{
			Code:    types.ParamErrorCode,
			Message: pkg.GetI18nMessage(c, types.ParamErrorCode),
		})
		return
	}
	todo := model.Todo{Title: input.Title, Description: input.Description}
	err := service.CreateTodo(todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, TodoRes{
		AppResponse: types.AppResponse{
			Code:    types.SucceedCode,
			Message: pkg.GetI18nMessage(c, types.SucceedCode),
		},
		Data: todo,
	})
}

// UpdateTodo godoc
// @Summary update a todo item
// @Schemes
// @Description update a todo item
// @Tags todo
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Router /todo/:id [put]
// @Success 200 {object} TodoRes
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Param   id     path    string     true        "ID"
// @Param   data     body    TodoInput     true        "data"
func (t TodoController) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.AppResponse{
			Code:    types.ParamErrorCode,
			Message: pkg.GetI18nMessage(c, types.ParamErrorCode),
		})
		return
	}
	input := UpdateTodo{
		ID: uid,
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, types.AppResponse{
			Code:    types.ParamErrorCode,
			Message: pkg.GetI18nMessage(c, types.ParamErrorCode),
		})
		return
	}
	todo := model.Todo{
		ID:          uid,
		Title:       input.Title,
		Description: input.Description,
	}
	if _, err := service.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, TodoRes{
		AppResponse: types.AppResponse{
			Code:    types.SucceedCode,
			Message: pkg.GetI18nMessage(c, types.SucceedCode),
		},
		Data: todo,
	})
}

// DeleteTodo godoc
// @Summary delete a todo item
// @Schemes
// @Description delete a todo item
// @Tags todo
// @Accept json
// @Produce json
// @security  ApiKeyAuth
// @Success 200 {object} types.AppResponse
// @Success 400 {object} types.AppResponse
// @Success 500 {object} types.AppResponse
// @Router /todo/:id [delete]
// @Param   id     path    string     true        "ID"
func (t TodoController) DeleteTodo(c *gin.Context) {
	var todo model.Todo
	var uid uint64
	id := c.Params.ByName("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.DeleteTodo(&todo, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.NewAppResponse(types.SystemErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, types.NewAppResponse(types.SucceedCode, id+" has been deleted"))
}
