// @title User API
// @version 1.0
// @description API for user management in Go
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

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
	_ "github.com/lib/pq"            // PostgreSQL driver
	"github.com/swaggo/http-swagger" // Swagger handler
	_ "go-crud-api/docs"             // Import the generated Swagger docs
	"go-crud-api/internal/controllers"
	"go-crud-api/internal/repos"
	"go-crud-api/internal/services"
	"go-crud-api/pkg/router"
)

// @title User API
// @version 1.0
// @description API for user management in Go
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /

func main() {
	db := initializeDatabase()
	defer closeDatabase(db)

	userController := initializeController(db)
	r := setupRouter(userController)
	startServer(r)
	handleGracefulShutdown()
}

func initializeDatabase() *sql.DB {
	dbURL := getDatabaseURL()
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	configureDBPooling(db)
	return db
}

func getDatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	return dbURL
}

func closeDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Error closing database: %v", err)
	}
}

func configureDBPooling(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func initializeController(db *sql.DB) *controllers.UserController {
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
