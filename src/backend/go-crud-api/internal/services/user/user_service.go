package user_service

import (
	usermodel "go-crud-api/internal/models/user"
	repository "go-crud-api/internal/repos/user"
)

type UserService interface {
	CreateUser(user *usermodel.User) error
	GetUserByID(id uint) (*usermodel.User, error)
	UpdateUser(user *usermodel.User) error
	DeleteUser(id uint) error
	GetUsers() ([]usermodel.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *usermodel.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*usermodel.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *usermodel.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) GetUsers() ([]usermodel.User, error) {
	return s.repo.GetUsers()
}
