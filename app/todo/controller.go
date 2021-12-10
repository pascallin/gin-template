package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/pkg"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/todo")

	router.GET("/", func(c *gin.Context) {
		var input Pagination
		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var todos []Todo
		err := getAllTodo(&todos, input.Page, input.PageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": todos})
			return
		}
	})
	router.POST("/", func(c *gin.Context) {
		var input CreateTodoInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo := Todo{Title: input.Title, Description: input.Description}
		err := createTodo(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": todo})
	})

	router.GET("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var todo Todo
		err := getTodo(&todo, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": todo})
		}
	})

	router.PUT("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		uid, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := UpdateTodoInput{
			ID: uid,
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo := Todo{GormModel: pkg.GormModel{ID: uid}, Title: input.Title, Description: input.Description}
		if err, _ := updateTodo(&todo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	router.DELETE("/:id", func(c *gin.Context) {
		var todo Todo
		var uid uint64
		id := c.Params.ByName("id")
		uid, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err, _ = deleteTodo(&todo, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": id + " has been deleted"})
	})
}
