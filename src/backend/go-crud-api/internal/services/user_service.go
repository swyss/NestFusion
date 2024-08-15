package services

import (
	"go-crud-api/internal/models"
	"go-crud-api/internal/repos"
)

type UserService struct {
	Repo *repos.UserRepository
}

func NewUserService(repo *repos.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(u *models.User) error {
	return s.Repo.CreateUser(u)
}

func (s *UserService) UpdateUser(u *models.User) error {
	return s.Repo.UpdateUser(u)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
