package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger" // Swagger handler
	_ "go-crud-api/docs"             // Import the generated Swagger docs
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/db"          // Separate package for database initialization
	"go-crud-api/internal/db/influxdb" // InfluxDB initialization
	"go-crud-api/internal/db/redisdb"  // Redis initialization
	"go-crud-api/internal/repos"
	"go-crud-api/internal/services"
	"go-crud-api/pkg/router"
)

func main() {
	// Initialize all services
	db := db.InitializePostgres()
	defer db.Close()

	redisClient := redisdb.InitializeRedis()
	defer redisClient.Close()

	influxClient := influxdb.InitializeInfluxDB()
	defer influxClient.Close()

	// Initialize controllers
	userController := initializeUserController(db)

	// Setup and start server
	r := setupRouter(userController)
	startServer(r)

	// Handle graceful shutdown
	handleGracefulShutdown()
}

func initializeUserController(db *sql.DB) *controllers.UserController {
	userRepo := repos.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	return controllers.NewUserController(userService)
}

func setupRouter(userController *controllers.UserController) *mux.Router {
	r := router.NewRouter(userController)

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}

func startServer(handler http.Handler) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router.JSONContentTypeMiddleware(handler),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :8000: %v", err)
		}
	}()
}

func handleGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server := &http.Server{}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
