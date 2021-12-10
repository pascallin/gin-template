package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/example")

	taskRouter := router.Group("/task")
	taskRouter.GET("/", func(c *gin.Context) {
		input := GetTaskListInput{}
		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err, tasks := getTasksData(findTasksCond{Title: input.Title}, input.Page, input.PageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": tasks})
	})
	taskRouter.POST("/", func(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"id": id})
	})
	taskRouter.GET("/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := getTaskById(id)
		if result == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})
	taskRouter.PUT("/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var task UpdateTaskInput
		c.BindJSON(&task)
		err, result := updateTaskData(id, &task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no task was found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})
	taskRouter.DELETE("/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = removeTaskData(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "okay"})
	})

	todoRouter := router.Group("/todo")
	todoRouter.GET("/", func(c *gin.Context) {
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
	todoRouter.POST("/", func(c *gin.Context) {
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
	todoRouter.GET("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var todo Todo
		err := getTodo(&todo, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": todo})
		}
	})
	todoRouter.PUT("/:id", func(c *gin.Context) {
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
	todoRouter.DELETE("/:id", func(c *gin.Context) {
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
