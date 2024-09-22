package user_router

import (
	"log"

	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/user/controllers"
	userrepo "go-crud-api/internal/user/repositories"
	userservice "go-crud-api/internal/user/services"
	database "go-crud-api/pkg/databases/postgres"
)

const (
	CreateUserRoute  = "/"
	GetUserRoute     = "/:id"
	RegisterRoute    = "/register"
	LoginRoute       = "/login"
	GetAllUsersRoute = "/all"
)

// initializeUserComponents initializes the user-related components such as repository, service, and controller.
func initializeUserComponents() (*controllers.UserController, error) {
	userRepo := userrepo.NewUserRepository(database.PostgresDB)
	userService := userservice.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	return userController, nil
}

// RegisterUserRoutes registers all user-related routes, including creating, getting users, and login/register.
func RegisterUserRoutes(r *gin.RouterGroup) {
	userController, err := initializeUserComponents()
	if err != nil {
		log.Fatalf("Failed to initialize user components: %v", err)
		return
	}

	// Register routes
	r.POST(CreateUserRoute, userController.CreateUser)  // Route for creating a user
	r.GET(GetUserRoute, userController.GetUser)         // Route for getting a user by ID
	r.POST(RegisterRoute, userController.RegisterUser)  // Route for user registration
	r.POST(LoginRoute, userController.LoginUser)        // Route for user login
	r.GET(GetAllUsersRoute, userController.GetAllUsers) // Route to get all users
}
