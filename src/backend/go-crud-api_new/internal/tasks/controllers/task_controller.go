package task_controllers

import (
	models "go-crud-api/internal/tasks/models"
	"go-crud-api/internal/tasks/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
}

var _taskService *task_services.TaskService

func NewTaskController() *TaskController {

	_taskService = task_services.NewTaskService()
	return &TaskController{}
}

func (controller *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := _taskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.IndentedJSON(http.StatusCreated, tasks)
}

func (controller *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := _taskService.GetAllTasks()
	if err != nil {
		// If there is an error, respond with a 500 status code.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Respond with the list of tasks and a 200 status code.
	c.IndentedJSON(http.StatusOK, tasks)
}

func (controllers *TaskController) UpdateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		// If decoding the request body fails, respond with a 400 status code.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task.ID = uint(id)
	tasks, err := _taskService.UpdateTask(&task)
	if err != nil {
		// If there is an error while updating the task, respond with a 500 status code.
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update task!"})
		return
	}
	// Respond with all the tasks and a 200 status code.
	c.IndentedJSON(http.StatusOK, tasks)
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	tasks, err := _taskService.DeleteTask(uint(id))
	// If there is an error while deleting the task, respond with a 500 status code.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
		return
	}

	// Respond with all the tasks and a 201 status code.
	c.IndentedJSON(200, tasks)
}

func (controller *TaskController) MarkTaskAsDone(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	tasks, err := _taskService.MarkTaskAsDone(uint(id))
	// If there is an error while deleting the task, respond with a 500 status code.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update task"})
		return
	}

	// Respond with all the tasks and a 201 status code.
	c.IndentedJSON(200, tasks)
}
