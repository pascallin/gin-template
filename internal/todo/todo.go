package todo

import (
	"github.com/pascallin/go-web/internal/common"
	databases "github.com/pascallin/go-web/internal/pkg/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTodoList(c *gin.Context) {
	var input common.Pagination
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todo []Todo
	err := getAllTodo(&todo, input.Page, input.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
		return
	}
}

func getATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	err := getTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func createATodo(c *gin.Context) {
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := Todo{Title:input.Title, Description:input.Description}
	err := createTodo(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func updateATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := UpdateTodoInput{
		ID:uid,
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := Todo{GormModel: databases.GormModel{ID: uid}, Title:input.Title, Description:input.Description}
	if err, _ := updateTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func deleteATodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	err, _ := deleteTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
