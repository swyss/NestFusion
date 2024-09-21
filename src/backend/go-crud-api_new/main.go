package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go-crud-api/pkg/router"
	"go-crud-api/pkg/startup"
	"go-crud-api/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	printAsciiArtTitle()

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Command-line flags
	resetDocker := flag.Bool("reset-docker", false, "Reset Docker and start environment")
	startAppOnly := flag.Bool("app-only", false, "Start the application without Docker setup")
	flag.Parse()

	// Select workflow based on flags
	if *resetDocker {
		utils.PrintWarning("Resetting Docker environment...\n")
		utils.StartSpinner(utils.FormatWarning, "Resetting Docker")
		startup.ResetDockerEnvironment()
		utils.StopSpinner()
		utils.PrintSuccess("Docker environment reset and started.\n")
		startApp()
	} else if *startAppOnly {
		utils.PrintSuccess("Starting application without Docker setup...\n")
		startApp()
	} else {
		if !startup.IsDockerRunning() {
			utils.PrintError("Docker environment not running. Starting Docker...\n")
			utils.StartSpinner(utils.FormatError, "Starting Docker")
			startup.StartDockerEnvironment()
			utils.StopSpinner()
		}
		utils.PrintInfo("Starting application after Docker setup...\n")
		startApp()
	}
}

// printAsciiArtTitle outputs the ASCII art title at startup
func printAsciiArtTitle() {
	fmt.Println(`
   _  __        __  ____         _         
  / |/ /__ ___ / /_/ __/_ _____ (_)__  ___ 
 /    / -_|_-</ __/ _// // (_-</ / _ \/ _ \
/_/|_/\__/___/\__/_/  \_,_/___/_/\___/_//_/
`)
}

// startApp handles service initialization and server startup
func startApp() {
	utils.PrintInfo("Initializing application services...\n")
	utils.StartSpinner(utils.FormatInfo, "Initializing Services")

	dbPostgres, redisClient, influxClient := startup.InitializeServices()

	defer func() {
		// Close all services on exit
		sqlDB, _ := dbPostgres.DB()
		if err := sqlDB.Close(); err != nil {
			return
		}
		if err := redisClient.Close(); err != nil {
			return
		}
		influxClient.Close()
		utils.PrintSuccess("All services stopped successfully.\n")
	}()

	// Stop spinner for services
	utils.StopSpinner()
	utils.PrintSuccess("✔ Services initialized!\n")

	// Start the HTTP server using Gin
	r := gin.Default()
	router.SetupRoutes(r)
	startServer(r)
	utils.PrintSuccess("App is running\n")
	handleGracefulShutdown()
}

// startServer starts the HTTP server
func startServer(router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		printAsciiArtTitle()
		utils.PrintSuccess("HTTP server running on port 8000...\n")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

// handleGracefulShutdown ensures app shuts down properly on interrupt
func handleGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	utils.PrintWarning("Graceful shutdown initiated...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Forced to shutdown: %v", err)
	}

	utils.PrintSuccess("Server exited.\n")
}
