package task_router

import (
	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/tasks/controllers"
)

func RegisterTaskRoutes(r *gin.Engine) {
	taskController := controllers.NewTaskController()

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("", taskController.GetAllTasks)
		taskRoutes.POST("", taskController.CreateTask)
		taskRoutes.PUT("/markAsDone/:id", taskController.MarkTaskAsDone)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
	}
}
