package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go-crud-api/internal/startup"
	"go-crud-api/pkg/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	controllers := startup.InitializeControllers(dbPostgres)

	// Setup and start the HTTP server
	r := setupRouter(controllers)
	startServer(r)

	// Handle graceful server shutdown on OS interrupt signals
	handleGracefulShutdown()
}

// setupRouter configures the HTTP router
func setupRouter(controllers *startup.Controllers) *gin.Engine {
	// Initialize router with all routes
	r := router.NewRouter(
		controllers.UserController,
		controllers.AuthController,
		controllers.RoleController,
		controllers.InfoController,
	)

	return r
}

// startServer starts the HTTP server in a new goroutine.
func startServer(handler *gin.Engine) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
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

	server := &http.Server{
		Addr: ":8000", // Server address must match the one used in startServer
	}

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown")
		log.Fatal(err)
	}
	log.Println("Server exiting")
}
