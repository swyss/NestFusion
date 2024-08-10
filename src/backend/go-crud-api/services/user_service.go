package services

import (
	"go-crud-api/models"
	"go-crud-api/repos"
)

type UserService struct {
	Repo *repos.UserRepo
}

func NewUserService(repo *repos.UserRepo) *UserService {
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
