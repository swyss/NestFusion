package repositories

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/gorm"
)

// UserRoleRepository handles the database operations for user roles.
type UserRoleRepository struct {
	DB *gorm.DB
}

// NewUserRoleRepository initializes a new UserRoleRepository.
func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{DB: db}
}

// GetAll retrieves all roles.
func (r *UserRoleRepository) GetAll() ([]models.UserRole, error) {
	var roles []models.UserRole
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Create creates a new role.
func (r *UserRoleRepository) Create(role *models.UserRole) error {
	return r.DB.Create(role).Error
}
