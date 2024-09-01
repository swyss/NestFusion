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
  
}
func TaskRouter(router *mux.Router, taskController *controllers.TaskController) *mux.Router{
	router.HandleFunc("/tasks", taskController.GetAllTasks).Methods("GET")
  router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
  router.HandleFunc("/tasks/markTaskAsDone/{id}", taskController.MarkTaskAsDone).Methods("PUT")
  router.HandleFunc("/tasks/{id}", taskController.UpdateTask).Methods("PUT")
  router.HandleFunc("/tasks/{id}", taskController.DeleteTask).Methods("DELETE")

	// Swagger UI route (falls noch nicht hinzugefügt)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
  return router
}

// JSONContentTypeMiddleware sets the Content-Type header to "application/json".
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
	return router
}
