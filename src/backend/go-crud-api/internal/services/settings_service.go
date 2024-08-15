package services

import (
	"go-crud-api/internal/models"
	"go-crud-api/internal/repos"
)

type SettingService struct {
	Repo *repos.SettingRepository
}

func NewSettingService(repo *repos.SettingRepository) *SettingService {
	return &SettingService{Repo: repo}
}

func (s *SettingService) Create(setting *models.Setting) error {
	return s.Repo.Create(setting)
}

func (s *SettingService) GetAll() ([]models.Setting, error) {
	return s.Repo.FindAll()
}
