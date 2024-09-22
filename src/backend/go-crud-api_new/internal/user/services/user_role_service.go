package services

import (
	models "go-crud-api/internal/user/models"
	"go-crud-api/internal/user/repositories"
)

// UserRoleService provides methods for role operations.
type UserRoleService struct {
	UserRoleRepo *repositories.UserRoleRepository
}

// NewUserRoleService initializes a new UserRoleService.
func NewUserRoleService(repo *repositories.UserRoleRepository) *UserRoleService {
	return &UserRoleService{
		UserRoleRepo: repo,
	}
}

// GetAllRoles retrieves all roles.
func (s *UserRoleService) GetAllRoles() ([]models.UserRole, error) {
	return s.UserRoleRepo.GetAll()
}

// CreateRole creates a new role.
func (s *UserRoleService) CreateRole(role *models.UserRole) error {
	return s.UserRoleRepo.Create(role)
}
