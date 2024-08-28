package startup

import (
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/repos"
	"go-crud-api/internal/services"
	"gorm.io/gorm"
	"log"
)

// InitializeControllers initializes all controllers and returns them.
func InitializeControllers(db *gorm.DB) *controllers.UserController {
	log.Println("Initializing controllers...")

	// Initialize UserRepository and UserService
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Initialize UserController
	userController := controllers.NewUserController(userService)
	log.Println("UserController initialized successfully")

  log.Println("Testing")
	return userController
}

func InitializeTaskController(db *gorm.DB) *controllers.TaskController {
	log.Println("Initializing controllers...")


  taskRepo := repos.NewTaskRepository(db)
  taskService := services.NewTaskService(taskRepo)

	// Initialize UserController
  taskController := controllers.NewTaskController(taskService)

	log.Println("TaskController initialized successfully")

	return taskController
}


