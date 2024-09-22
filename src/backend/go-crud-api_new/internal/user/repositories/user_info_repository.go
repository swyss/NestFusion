package repositories

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/gorm"
)

// UserInfoRepository handles the database operations for user info.
type UserInfoRepository struct {
	DB *gorm.DB
}

// NewUserInfoRepository initializes a new UserInfoRepository.
func NewUserInfoRepository(db *gorm.DB) *UserInfoRepository {
	return &UserInfoRepository{DB: db}
}

// GetAll retrieves all user info records.
func (r *UserInfoRepository) GetAll() ([]models.UserInfo, error) {
	var userInfo []models.UserInfo
	if err := r.DB.Find(&userInfo).Error; err != nil {
		return nil, err
	}
	return userInfo, nil
}

// Create creates a new user info record.
func (r *UserInfoRepository) Create(userInfo *models.UserInfo) error {
	return r.DB.Create(userInfo).Error
}
