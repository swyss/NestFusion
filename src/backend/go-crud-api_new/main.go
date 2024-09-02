package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go-crud-api/pkg/router"
	"go-crud-api/pkg/startup"
)

func main() {
	log.Println("Starting application...")

	// Initialize services
	log.Println("Initializing services...")
	dbPostgres, redisClient, influxClient := startup.InitializeServices()
	defer func() {
		sqlDB, err := dbPostgres.DB()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		log.Println("Database connection closed successfully.")
	}()
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Fatalf("Failed to close Redis client: %v", err)
		}
		log.Println("Redis client closed successfully.")
	}()
	defer func() {
		influxClient.Close()
		log.Println("InfluxDB client closed successfully.")
	}()

	log.Println("All services initialized successfully.")

	// Setup and start the HTTP server using Gin
	r := gin.Default()
	router.SetupRoutes(r)

	startServer(r)

	// Handle graceful server shutdown on OS interrupt signals
	handleGracefulShutdown()
}

// startServer starts the HTTP server in a new goroutine.
func startServer(router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		log.Println("Starting HTTP server on port 8000...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :8000: %v", err)
		}
	}()
	log.Println("Server started on port 8000.")
}

// handleGracefulShutdown manages graceful server shutdown upon receiving OS signals.
func handleGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit // Block until a signal is received
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully.")
}
