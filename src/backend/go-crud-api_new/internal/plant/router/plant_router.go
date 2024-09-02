package plant_router

import (
	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/plant/controllers"
	repositories "go-crud-api/internal/plant/repositories"
	services "go-crud-api/internal/plant/services"
	"go-crud-api/pkg/databases/postgres"
)

func initializePlantComponents() (*controllers.PlantController, error) {
	repo := repositories.NewPlantRepository(postgres.PostgresDB)
	service := services.NewPlantService(repo)
	controller := controllers.NewPlantController(service)
	return controller, nil
}

func registerPlantHandlers(r *gin.RouterGroup, controller *controllers.PlantController) {
	r.POST("/", controller.CreatePlant)
	r.GET("/:id", controller.GetPlant)
	// Additional plant routes can be added here
}

func RegisterPlantRoutes(r *gin.RouterGroup) {
	controller, err := initializePlantComponents()
	if err != nil {
		// Handle initialization error (optional)
		return
	}
	registerPlantHandlers(r, controller)
}
