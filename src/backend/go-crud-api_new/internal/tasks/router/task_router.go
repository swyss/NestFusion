package task_router

import (
	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/tasks/controllers"
	repositories "go-crud-api/internal/tasks/repositories"
	services "go-crud-api/internal/tasks/services"
	database "go-crud-api/pkg/databases/postgres"
)

func RegisterTaskRoutes(r *gin.RouterGroup) {
	// Initialize the repository, service, and controller
	taskRepo := repositories.NewTaskRepository(database.PostgresDB)
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	// Register routes
	r.POST("/", taskController.CreateTask)
	r.GET("/:id", taskController.GetTask)
	// Additional task routes can be added here
}
