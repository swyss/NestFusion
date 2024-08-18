package init

import (
	"database/sql"
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/logger"
	"go-crud-api/internal/repos"
	"go-crud-api/internal/services"
)

// InitializeControllers initializes all controllers and returns them.
func InitializeControllers(db *sql.DB, log *logger.Logger) *controllers.UserController {
	log.InfoMsg("Initializing controllers...")

	// Initialize UserRepository and UserService
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Initialize UserController
	userController := controllers.NewUserController(userService)
	log.InfoMsg("UserController initialized successfully")

	return userController
}
