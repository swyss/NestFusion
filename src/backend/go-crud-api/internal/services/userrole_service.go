package services

import (
	"go-crud-api/internal/models"
	"go-crud-api/internal/repos"
)

type UserRoleService struct {
	Repo *repos.UserRoleRepository
}

func NewUserRoleService(repo *repos.UserRoleRepository) *UserRoleService {
	return &UserRoleService{Repo: repo}
}

func (s *UserRoleService) Create(role *models.UserRole) error {
	return s.Repo.Create(role)
}

func (s *UserRoleService) GetAll() ([]models.UserRole, error) {
	return s.Repo.FindAll()
}
