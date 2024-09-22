package user_services

import (
	models "go-crud-api/internal/user/models"
	repositories "go-crud-api/internal/user/repositories"
)

type UserInfoService struct {
	repository *repositories.UserInfoRepository
}

func NewUserInfoService(repo *repositories.UserInfoRepository) *UserInfoService {
	return &UserInfoService{repository: repo}
}

// CreateUserInfo creates new user info in the database.
func (service *UserInfoService) CreateUserInfo(info *models.UserInfo) error {
	return service.repository.CreateUserInfo(info)
}
