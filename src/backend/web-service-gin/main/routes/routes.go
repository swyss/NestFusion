// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/controllers"
)

func RegisterRoutes(router *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {
	// User routes
	router.POST("/create-user", userController.CreateUser)
	router.GET("/get-user", userController.GetUser)
	router.GET("/get-all-users", userController.GetAllUsers)
	router.PUT("/update-user", userController.UpdateUser)
	router.DELETE("/delete-user", userController.DeleteUser)

	// Task routes
	//router.POST("/create-task", taskController.CreateTask)
	//router.GET("/get-task", taskController.GetTask)
	//router.GET("/get-all-tasks", taskController.GetAllTasks)
	//router.PUT("/update-task", taskController.UpdateTask)
	//router.DELETE("/delete-task", taskController.DeleteTask)
}
