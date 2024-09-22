package user_repositories

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/gorm"
)

type UserInfoRepository struct {
	db *gorm.DB
}

func NewUserInfoRepository(db *gorm.DB) *UserInfoRepository {
	return &UserInfoRepository{db: db}
}

// CreateUserInfo inserts new user info into the database.
func (repo *UserInfoRepository) CreateUserInfo(info *models.UserInfo) error {
	return repo.db.Create(info).Error
}
