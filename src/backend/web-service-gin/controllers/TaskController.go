package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/services"
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (controller *TaskController) GetTasks(context *gin.Context) {
	taskService := services.NewTaskService()
	tasks := taskService.GetTasks()

	context.IndentedJSON(200, tasks)
}

func (controller *TaskController) GetTaskById(c *gin.Context) {
	service := services.NewTaskService()

	id := c.Param("id")
	task, err := service.GetTaskById(id)
	if err != nil {
		// Task not found
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(200, task)
}
