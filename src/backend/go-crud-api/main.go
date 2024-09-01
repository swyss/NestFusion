package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go-crud-api/internal/startup"
	"go-crud-api/pkg/router"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Starting application...")

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Fatal(err)
		return
	}
	log.Println(".env file loaded successfully")

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

	userController := startup.InitializeControllers(dbPostgres)
	taskController := startup.InitializeTaskController(dbPostgres)

	r := setupRouter(userController, taskController)

	startServer(r)
	handleGracefulShutdown()
}

func setupRouter(userController *startup.Controllers, taskController http.Handler) *gin.Engine {
	r := router.NewRouter(
		userController.UserController,
		userController.AuthController,
		userController.RoleController,
		userController.InfoController,
	)

	// TaskController als http.Handler in den Gin-Router einbinden
	r.Any("/tasks/*any", gin.WrapH(taskController))

	return r
}

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

func handleGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server := &http.Server{
		Addr: ":8000",
	}

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown")
		log.Fatal(err)
	}
	log.Println("Server exiting")
}
