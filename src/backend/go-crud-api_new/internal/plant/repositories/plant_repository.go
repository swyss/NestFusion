package plant_repositories

import (
	"go-crud-api/internal/plant/models"
	"gorm.io/gorm"
)

type PlantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) *PlantRepository {
	return &PlantRepository{db: db}
}

func (repository *PlantRepository) CreatePlant(plant *models.Plant) error {
	return repository.db.Create(plant).Error
}

func (repository *PlantRepository) GetPlantByID(id uint) (*models.Plant, error) {
	var plant models.Plant
	err := repository.db.First(&plant, id).Error
	return handleGetResult(&plant, err)
}

func handleGetResult(plant *models.Plant, err error) (*models.Plant, error) {
	if err != nil {
		return nil, err
	}
	return plant, nil
}
