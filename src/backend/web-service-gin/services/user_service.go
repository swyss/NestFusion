package services

import (
	"github.com/google/uuid"
	"web-service-gin/models"
	"web-service-gin/repositories"
	_ "web-service-gin/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) CreateUser(username, email, password string) (*models.User, error) {
	user := &models.User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: password,
	}
	err := service.Repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	return service.Repo.FindById(id)
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.Repo.FindAll()
}

func (service *UserService) UpdateUser(id, username, email, password string) (*models.User, error) {
	user, err := service.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	user.Username = username
	user.Email = email
	user.Password = password
	err = service.Repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) DeleteUser(id string) error {
	return service.Repo.Delete(id)
}
