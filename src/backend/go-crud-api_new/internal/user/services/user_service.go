package user_services

import (
	models "go-crud-api/internal/user/models"
	repositories "go-crud-api/internal/user/repositories"
)

// UserService provides user-related services.
type UserService struct {
	repository *repositories.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

// RegisterUser registers a new user.
func (service *UserService) RegisterUser(user *models.User) error {
	return service.repository.CreateUser(user)
}

// GetUser retrieves a user by ID.
func (service *UserService) GetUser(userID uint) (*models.User, error) {
	return service.repository.GetUserByID(userID)
}
