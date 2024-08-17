package services

import (
	"go-crud-api/internal/models"
	"go-crud-api/internal/repos"
)

// UserServiceInterface defines the methods that the UserService must implement.
type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(u *models.User) error
	UpdateUser(u *models.User) error
	DeleteUser(id int) error
}

// UserService provides business logic for user-related operations.
type UserService struct {
	Repo *repos.UserRepository
}

// NewUserService creates a new instance of UserService with the provided UserRepository.
func NewUserService(repo *repos.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Ensure UserService implements UserServiceInterface.
var _ UserServiceInterface = &UserService{}

// GetAllUsers retrieves all users by delegating the call to the UserRepository.
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

// GetUserByID retrieves a user by their ID by delegating the call to the UserRepository.
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

// CreateUser creates a new user by delegating the call to the UserRepository.
func (s *UserService) CreateUser(u *models.User) error {
	return s.Repo.CreateUser(u)
}

// UpdateUser updates an existing user by delegating the call to the UserRepository.
func (s *UserService) UpdateUser(u *models.User) error {
	return s.Repo.UpdateUser(u)
}

// DeleteUser deletes a user by their ID by delegating the call to the UserRepository.
func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
