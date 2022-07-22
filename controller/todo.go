package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/model"
	"github.com/pascallin/gin-template/service"
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

// GetTodos godoc
// @Summary get todo list
// @Description get todo list
// @Tags todo
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {array} Todo
// @Router /todo [get]
// @Param   page     query    number     true        "page number"
// @Param   pageSize    query    number     true        "page size"
func (t TodoController) GetTodos(c *gin.Context) {
	var input Pagination
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todos []model.Todo
	err := service.GetAllTodo(&todos, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": todos})
		return
	}
}

// GetTodo godoc
// @Summary get todo
// @Description get todo
// @Tags todo
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Success 200 {object} Todo
// @Router /todo/:id [get]
// @Param   id     path    string     true        "ID"
func (t TodoController) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo model.Todo
	err := service.GetTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": todo})
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
// @Router /todo [post]
// @Param   data     body    TodoInput     true        "data"
func (t TodoController) CreateTodo(c *gin.Context) {
	var input TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := model.Todo{Title: input.Title, Description: input.Description}
	err := service.CreateTodo(todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
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
// @Param   id     path    string     true        "ID"
// @Param   data     body    TodoInput     true        "data"
func (t TodoController) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := UpdateTodo{
		ID: uid,
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := model.Todo{
		ID:          uid,
		Title:       input.Title,
		Description: input.Description,
	}
	if _, err := service.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary delete a todo item
// @Schemes
// @Description delete a todo item
// @Tags todo
// @Accept json
// @Produce json
// @security  ApiKeyAuth
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id + " has been deleted"})
}
