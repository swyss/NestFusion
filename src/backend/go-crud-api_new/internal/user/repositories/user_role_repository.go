package user_repositories

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{db: db}
}

// CreateUserRole inserts a new role into the database.
func (repo *UserRoleRepository) CreateUserRole(role *models.UserRole) error {
	return repo.db.Create(role).Error
}
