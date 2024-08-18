package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/init"
	"go-crud-api/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger" // Swagger handler
	"go-crud-api/pkg/router"
)

func main() {
	log := logger.NewLogger()

	log.InfoMsg("Starting application...")

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.ErrorMsg(err)
		log.InfoMsg("Error loading .env file")
		return
	}
	log.InfoMsg(".env file loaded successfully")

	// Initialize services
	dbPostgres, redisClient, influxClient := init.InitializeServices(log)
	defer func(dbPostgres *sql.DB) {
		err := dbPostgres.Close()
		if err != nil {

		}
	}(dbPostgres)
	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {

		}
	}(redisClient)
	defer influxClient.Close()

	log.InfoMsg("All services initialized successfully")

	// Initialize controllers
	userController := init.InitializeControllers(dbPostgres, log)

	// Setup and start the HTTP server
	r := setupRouter(userController)
	startServer(r, log)

	// Handle graceful server shutdown on OS interrupt signals
	handleGracefulShutdown(log)
}

// setupRouter configures the HTTP router and adds Swagger UI.
func setupRouter(userController *controllers.UserController) *mux.Router {
	// Initialize router with user routes and attach Swagger handler
	r := router.NewRouter(userController)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}

// startServer starts the HTTP server in a new goroutine.
func startServer(handler http.Handler, log *logger.Logger) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router.JSONContentTypeMiddleware(handler),
	}

	// Start the server and log any errors except when the server is closed
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.ErrorMsg(err)
			log.InfoMsg("Could not listen on :8000")
		}
	}()
	log.InfoMsg("Server started on port 8000")
}

// handleGracefulShutdown manages graceful server shutdown upon receiving OS signals.
func handleGracefulShutdown(log *logger.Logger) {
	// Create a channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit // Block until a signal is received
	log.InfoMsg("Shutting down server...")

	// Create a context with a 5-second timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.ErrorMsg(err)
		log.InfoMsg("Server forced to shutdown")
	}
	log.InfoMsg("Server exiting")
}
