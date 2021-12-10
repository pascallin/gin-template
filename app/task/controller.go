package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/task")

	router.GET("/", func(c *gin.Context) {
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

	router.POST("/", func(c *gin.Context) {
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

	router.GET("/:id", func(c *gin.Context) {
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

	router.PUT("/:id", func(c *gin.Context) {
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

	router.DELETE("/:id", func(c *gin.Context) {
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
}
