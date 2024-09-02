package plant_services

import (
	"go-crud-api/internal/plant/models"
	repositories "go-crud-api/internal/plant/repositories"
)

// PlantService handles business logic for plants.
type PlantService struct {
	repo *repositories.PlantRepository
}

// NewPlantService initializes a new PlantService with the provided repository.
func NewPlantService(repository *repositories.PlantRepository) *PlantService {
	return &PlantService{repo: repository}
}

// RegisterPlant registers a new plant in the system.
func (s *PlantService) RegisterPlant(plant *models.Plant) error {
	return s.repo.CreatePlant(plant)
}

// FetchPlant retrieves a plant by its ID.
func (s *PlantService) FetchPlant(plantID uint) (*models.Plant, error) {
	return s.repo.GetPlantByID(plantID)
}

// GetPlant gets a plant summary by its ID.
func (s *PlantService) GetPlant(plantID uint) (interface{}, error) {
	plant, err := s.FetchPlant(plantID)
	if err != nil {
		return nil, err
	}
	// Return a simple summary of the plant.
	summary := map[string]interface{}{
		"id":          plant.ID,
		"name":        plant.Name,
		"description": plant.Description,
	}
	return summary, nil
}
