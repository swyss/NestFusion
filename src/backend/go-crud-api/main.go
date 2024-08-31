package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/startup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go-crud-api/pkg/router"
	"gorm.io/gorm"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {

	log.Println("Starting application...")

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Fatal(err)
		return
	}
	log.Println(".env file loaded successfully")

	// Initialize services
	dbPostgres, redisClient, influxClient := startup.InitializeServices()
	defer func(dbPostgres *gorm.DB) {
		sqlDB, err := dbPostgres.DB()
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatal(err)
		}
	}(dbPostgres)
	defer func(redisClient *redis.Client) {
		if err := redisClient.Close(); err != nil {
			log.Fatal(err)
		}
	}(redisClient)
	defer influxClient.Close()

	log.Println("All services initialized successfully")

	// Initialize controllers
	userController := startup.InitializeControllers(dbPostgres)
	taskController := startup.InitializeTaskController(dbPostgres)

	// Setup and start the HTTP server
	r := setupRouter(userController, taskController)
	startServer(r)

	// Handle graceful server shutdown on OS interrupt signals
	handleGracefulShutdown()
}

// setupRouter configures the HTTP router
func setupRouter(userController *controllers.UserController, taskController *controllers.TaskController) *mux.Router {
	// Initialize router with user routes
	r := router.NewRouter(userController)
	r = router.TaskRouter(r, taskController)

	// Swagger UI route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}

// CORSHandler sets the necessary headers for CORS.
func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// startServer starts the HTTP server in a new goroutine.
func startServer(handler http.Handler) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: CORSHandler(router.JSONContentTypeMiddleware(handler)), // Added CORSHandler here
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Could not listen on :8000")
			log.Fatal(err)
		}
	}()
	log.Println("Server started on port 8000")
}

// handleGracefulShutdown manages graceful server shutdown upon receiving OS signals.
func handleGracefulShutdown() {
	// Create a channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit // Block until a signal is received
	log.Println("Shutting down server...")

	// Create a context with a 5-second timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown")
		log.Fatal(err)
	}
	log.Println("Server exiting")
}
