package user_router

import (
	"go-crud-api/internal/user/controllers"
	"log"

	"github.com/gin-gonic/gin"
	usercontrollers "go-crud-api/internal/user/controllers"
	userrepo "go-crud-api/internal/user/repositories"
	userservice "go-crud-api/internal/user/services"
	database "go-crud-api/pkg/databases/postgres"
)

const (
	CreateUserRoute     = "/"
	GetUserRoute        = "/:id"
	RegisterRoute       = "/register"
	LoginRoute          = "/login"
	GetAllUsersRoute    = "/all"
	GetAllUserInfoRoute = "/info"
	CreateUserInfoRoute = "/info"
	GetAllRolesRoute    = "/roles"
	CreateRoleRoute     = "/roles"
)

// initializeUserComponents initializes the user-related components such as repository, service, and controller.
func initializeUserComponents() (*controllers.UserController, *controllers.UserInfoController, *controllers.UserRoleController, error) {
	// Initialize User components
	userRepo := userrepo.NewUserRepository(database.PostgresDB)
	userService := userservice.NewUserService(userRepo)
	userController := usercontrollers.NewUserController(userService)

	// Initialize UserInfo components
	userInfoRepo := userrepo.NewUserInfoRepository(database.PostgresDB)
	userInfoService := userservice.NewUserInfoService(userInfoRepo)
	userInfoController := usercontrollers.UserInfoController{UserInfoService: *userInfoService}

	// Initialize UserRole components
	userRoleRepo := userrepo.NewUserRoleRepository(database.PostgresDB)
	userRoleService := userservice.NewUserRoleService(userRoleRepo)
	userRoleController := usercontrollers.UserRoleController{UserRoleService: *userRoleService}

	return userController, &userInfoController, &userRoleController, nil
}

// RegisterUserRoutes registers all user-related routes, including creating, getting users, and login/register.
func RegisterUserRoutes(r *gin.RouterGroup) {
	userController, userInfoController, userRoleController, err := initializeUserComponents()
	if err != nil {
		log.Fatalf("Failed to initialize user components: %v", err)
		return
	}

	// Register user routes
	r.POST(CreateUserRoute, userController.CreateUser)  // Route for creating a user
	r.GET(GetUserRoute, userController.GetUser)         // Route for getting a user by ID
	r.POST(RegisterRoute, userController.RegisterUser)  // Route for user registration
	r.POST(LoginRoute, userController.LoginUser)        // Route for user login
	r.GET(GetAllUsersRoute, userController.GetAllUsers) // Route to get all users

	// Register user info routes
	r.GET(GetAllUserInfoRoute, userInfoController.GetAllUserInfo)  // Route to get all user info
	r.POST(CreateUserInfoRoute, userInfoController.CreateUserInfo) // Route to create user info

	// Register user role routes
	r.GET(GetAllRolesRoute, userRoleController.GetAllRoles) // Route to get all roles
	r.POST(CreateRoleRoute, userRoleController.CreateRole)  // Route to create a new role
}
