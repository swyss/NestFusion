package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
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
	dbPostgres := db.InitializePostgres()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(dbPostgres)

	redisClient := redisdb.InitializeRedis()
	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {

		}
	}(redisClient)

	influxClient := influxdb.InitializeInfluxDB()
	defer influxClient.Close()

	// Initialize controllers
	userController := initializeUserController(dbPostgres)
	userRoleController := initializeUserRoleController(redisClient)
	settingController := initializeSettingController(redisClient)

	// Setup and start server
	r := setupRouter(userController, userRoleController, settingController)
	startServer(r)

	// Handle graceful shutdown
	handleGracefulShutdown()
}

func initializeUserController(db *sql.DB) *controllers.UserController {
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return controllers.NewUserController(userService)
}

func initializeUserRoleController(redisClient *redis.Client) *controllers.UserRoleController {
	userRoleRepo := repos.NewUserRoleRepository(redisClient)
	userRoleService := services.NewUserRoleService(userRoleRepo)
	return controllers.NewUserRoleController(userRoleService)
}

func initializeSettingController(redisClient *redis.Client) *controllers.SettingController {
	settingRepo := repos.NewSettingRepository(redisClient)
	settingService := services.NewSettingService(settingRepo)
	return controllers.NewSettingController(settingService)
}

func setupRouter(userController *controllers.UserController, userRoleController *controllers.UserRoleController, settingController *controllers.SettingController) *mux.Router {
	r := router.NewRouter(userController, userRoleController, settingController)

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
