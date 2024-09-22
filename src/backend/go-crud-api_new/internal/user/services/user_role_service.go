package user_services

import (
	models "go-crud-api/internal/user/models"
	repositories "go-crud-api/internal/user/repositories"
)

type UserRoleService struct {
	repository *repositories.UserRoleRepository
}

func NewUserRoleService(repo *repositories.UserRoleRepository) *UserRoleService {
	return &UserRoleService{repository: repo}
}

// CreateUserRole creates a new role in the database.
func (service *UserRoleService) CreateUserRole(role *models.UserRole) error {
	return service.repository.CreateUserRole(role)
}
