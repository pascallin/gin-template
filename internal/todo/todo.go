package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodoList(c *gin.Context) {
	var todo []Todo
	err := GetAllTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func createATodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	err := CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func getATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	err := GetTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func updateATodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	err := GetTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func deleteATodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	err := DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
