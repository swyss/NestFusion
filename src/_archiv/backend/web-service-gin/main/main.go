package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"web-service-gin/controllers"
	"web-service-gin/main/routes"
	"web-service-gin/repositories"
	_ "web-service-gin/repositories"
	"web-service-gin/services"
)

func main() {
	// Redis-Client konfigurieren
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// User-Repository, -Service und -Controller initialisieren
	userRepo := repositories.NewUserRepository(rdb)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Task-Repository, -Service und -Controller initialisieren
	//taskRepo := repositories.NewTaskRepository(rdb)
	//taskService := services.NewTaskService(taskRepo)
	//taskController := controllers.NewTaskController(taskService)

	// Gin-Router konfigurieren und Routen registrieren
	//router := routes.SetupRouter(userController, taskController)
	router := routes.SetupRouter(userController, nil)

	// Beispiel-Route für Alben
	router.GET("/albums", getAlbums)

	// Server starten
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
