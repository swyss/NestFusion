package task_router

import (
	controllers "go-crud-api/internal/tasks/controllers"
	task_repositories "go-crud-api/internal/tasks/repositories"
	task_services "go-crud-api/internal/tasks/services"
	database "go-crud-api/pkg/databases/postgres"

	"github.com/gin-gonic/gin"
)

func initializeTaskComponents() *controllers.TaskController {
	taskRepo := task_repositories.NewTaskRepository(database.PostgresDB)
	taskService := task_services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)
	return taskController
}

func RegisterTaskRoutes(taskGroup *gin.RouterGroup) {
	taskController := initializeTaskComponents()

	{
		taskGroup.GET("", taskController.GetAllTasks)
		taskGroup.POST("", taskController.CreateTask)
		taskGroup.PUT("/markAsDone/:id", taskController.MarkTaskAsDone)
		taskGroup.DELETE("/:id", taskController.DeleteTask)
		taskGroup.PUT("/:id", taskController.UpdateTask)
	}
}
