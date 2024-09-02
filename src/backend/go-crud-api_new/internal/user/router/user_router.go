package user_router

import (
	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/user/controllers"
	userrepo "go-crud-api/internal/user/repositories"
	userservice "go-crud-api/internal/user/services"
	database "go-crud-api/pkg/databases/postgres"
)

const (
	CreateUserRoute = "/"
	GetUserRoute    = "/:id"
)

func initializeUserComponents() (*controllers.UserController, error) {
	userRepo := userrepo.NewUserRepository(database.PostgresDB)
	userService := userservice.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	return userController, nil
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	userController, err := initializeUserComponents()
	if err != nil {
		return
	}

	// Register routes
	r.POST(CreateUserRoute, userController.CreateUser)
	r.GET(GetUserRoute, userController.GetUser)
	// Additional user routes can be added here
}
