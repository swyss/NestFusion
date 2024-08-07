package routes

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/controllers"
)

// SetupRouter initializes the Gin router and registers the routes
func SetupRouter(userController *controllers.UserController, taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	// Register user routes
	RegisterRoutes(r, userController, taskController)

	return r
}
