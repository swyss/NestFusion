package services

import (
	models "go-crud-api/internal/user/models"
	"go-crud-api/internal/user/repositories"
)

// UserInfoService provides methods for user info operations.
type UserInfoService struct {
	UserInfoRepo *repositories.UserInfoRepository
}

// NewUserInfoService initializes a new UserInfoService.
func NewUserInfoService(repo *repositories.UserInfoRepository) *UserInfoService {
	return &UserInfoService{
		UserInfoRepo: repo,
	}
}

// GetAllUserInfo retrieves all user info records.
func (s *UserInfoService) GetAllUserInfo() ([]models.UserInfo, error) {
	return s.UserInfoRepo.GetAll()
}

// CreateUserInfo creates a new user info record.
func (s *UserInfoService) CreateUserInfo(userInfo *models.UserInfo) error {
	return s.UserInfoRepo.Create(userInfo)
}
