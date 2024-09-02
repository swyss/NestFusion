package plant_controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud-api/internal/plant/models"
	services "go-crud-api/internal/plant/services"
	"net/http"
	"strconv"
)

type PlantController struct {
	service *services.PlantService
}

func NewPlantController(service *services.PlantService) *PlantController {
	return &PlantController{service: service}
}

func (controller *PlantController) CreatePlant(c *gin.Context) {
	var plant models.Plant
	if err := c.ShouldBindJSON(&plant); err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := controller.service.RegisterPlant(&plant); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create plant")
		return
	}
	c.JSON(http.StatusCreated, plant)
}

func (controller *PlantController) GetPlant(c *gin.Context) {
	plantID, err := parsePlantID(c)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid plant ID")
		return
	}
	plant, err := controller.service.GetPlant(plantID)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Plant not found")
		return
	}
	c.JSON(http.StatusOK, plant)
}

func parsePlantID(c *gin.Context) (uint, error) {
	plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(plantID), nil
}

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}
