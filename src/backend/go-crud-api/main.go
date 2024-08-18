package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger" // Swagger handler
	_ "go-crud-api/docs"             // Import the generated Swagger docs
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/databases"          // Separate package for database initialization
	"go-crud-api/internal/databases/influxdb" // InfluxDB initialization
	"go-crud-api/internal/databases/redisdb"  // Redis initialization
	"go-crud-api/internal/repos"
	"go-crud-api/internal/services"
	"go-crud-api/pkg/router"
)

func main() {
	log.Println("Starting application...")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println(".env file loaded successfully")

	// Initialize the PostgreSQL database
	dbPostgres := databases.InitializePostgres()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing PostgreSQL connection: %v", err)
		}
	}(dbPostgres)

	// Initialize Redis client
	redisClient := redisdb.InitializeRedis()
	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}(redisClient)

	// Initialize InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	defer influxClient.Close()

	log.Println("All services initialized successfully")

	// Initialize user controller
	userController := initializeUserController(dbPostgres)

	// Setup and start the HTTP server
	r := setupRouter(userController)
	startServer(r)

	// Handle graceful server shutdown on OS interrupt signals
	handleGracefulShutdown()
}

// initializeUserController sets up the user controller with the necessary dependencies.
func initializeUserController(db *sql.DB) *controllers.UserController {
	// Create UserRepository and UserService, then return UserController
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return controllers.NewUserController(userService)
}

// setupRouter configures the HTTP router and adds Swagger UI.
func setupRouter(userController *controllers.UserController) *mux.Router {
	// Initialize router with user routes and attach Swagger handler
	r := router.NewRouter(userController)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}

// startServer starts the HTTP server in a new goroutine.
func startServer(handler http.Handler) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router.JSONContentTypeMiddleware(handler),
	}

	// Start the server and log any errors except when the server is closed
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :8000: %v", err)
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

	// Attempt to gracefully shutdown the server
	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
