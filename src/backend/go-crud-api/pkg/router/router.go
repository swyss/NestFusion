package router

import (
	"github.com/gin-gonic/gin"
	usercontrollers "go-crud-api/internal/controllers/user"
)

// NewRouter configures the routing for users, user roles, settings, and other controllers.
func NewRouter(
	userController *usercontrollers.UserController,
	authController *usercontrollers.AuthController,
	roleController *usercontrollers.RoleController,
	infoController *usercontrollers.InfoController,
) *gin.Engine {
	router := gin.Default()

	// User routes
	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Authentication route
	router.POST("/auth", authController.Authenticate)

	// Role assignment route
	router.POST("/users/:id/assign-role", roleController.AssignRole)

	// User info setting route
	router.POST("/users/:id/set-userinfo", infoController.SetUserInfo)

	return router
}
