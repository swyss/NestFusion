package startup

import (
	"log"

	usercontroller "go-crud-api/internal/controllers/user"
	userrepo "go-crud-api/internal/repos/user"
	userservice "go-crud-api/internal/services/user"
	"gorm.io/gorm"
)

type Controllers struct {
	UserController *usercontroller.UserController
	AuthController *usercontroller.AuthController
	RoleController *usercontroller.RoleController
	InfoController *usercontroller.InfoController
}

// InitializeControllers initializes all controllers and returns them.
func InitializeControllers(db *gorm.DB) *Controllers {
	log.Println("Initializing controllers...")

	// Initialize repositories
	userRepo := userrepo.NewUserRepository(db)
	authRepo := userrepo.NewAuthRepository(db)
	roleRepo := userrepo.NewRoleRepository(db)
	infoRepo := userrepo.NewInfoRepository(db)

	// Initialize services
	userService := userservice.NewUserService(userRepo)
	authService := userservice.NewAuthService(authRepo)
	roleService := userservice.NewRoleService(roleRepo)
	infoService := userservice.NewInfoService(infoRepo)

	// Initialize controllers with the required services
	userController := usercontroller.NewUserController(userService)
	authController := usercontroller.NewAuthController(authService)
	roleController := usercontroller.NewRoleController(roleService)
	infoController := usercontroller.NewInfoController(infoService)

	log.Println("All controllers initialized successfully")

	return &Controllers{
		UserController: userController,
		AuthController: authController,
		RoleController: roleController,
		InfoController: infoController,
	}
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


