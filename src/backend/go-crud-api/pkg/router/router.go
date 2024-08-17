package router

import (
	"go-crud-api/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configures the routing for users, user roles, and settings.
func NewRouter(userController *controllers.UserController) *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}

// JSONContentTypeMiddleware sets the Content-Type header to "application/json".
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
