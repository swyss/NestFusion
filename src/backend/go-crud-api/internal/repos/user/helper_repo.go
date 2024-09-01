package user_repository

import (
	usermodel "go-crud-api/internal/models/user"
	"gorm.io/gorm"
)

// AuthRepository defines the methods for the authentication repository.
type AuthRepository interface {
	FindByUserID(userID uint) (*usermodel.AuthInput, error)
}

// GormAuthRepository is the implementation of AuthRepository using GORM.
type GormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &GormAuthRepository{db}
}

func (r *GormAuthRepository) FindByUserID(userID uint) (*usermodel.AuthInput, error) {
	var auth usermodel.AuthInput
	if err := r.db.Where("user_id = ?", userID).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

// RoleRepository defines the methods for the role repository.
type RoleRepository interface {
	AssignRole(userID, roleID uint) error
}

// GormRoleRepository is the implementation of RoleRepository using GORM.
type GormRoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &GormRoleRepository{db}
}

func (r *GormRoleRepository) AssignRole(userID, roleID uint) error {
	// This example assumes there is a field `RoleID` in the User model.
	// If roles are managed via a join table, the implementation will be different.
	err := r.db.Model(&usermodel.User{}).Where("id = ?", userID).Update("role_id", roleID).Error
	if err != nil {
		return err
	}
	return nil
}

// InfoRepository defines the methods for the user info repository.
type InfoRepository interface {
	SetUserInfo(userID, infoID uint) error
}

// GormInfoRepository is the implementation of InfoRepository using GORM.
type GormInfoRepository struct {
	db *gorm.DB
}

func NewInfoRepository(db *gorm.DB) InfoRepository {
	return &GormInfoRepository{db}
}

func (r *GormInfoRepository) SetUserInfo(userID, infoID uint) error {
	// This example assumes there is a separate table for user info linked by user_id.
	// If the information is stored directly on the user table, the implementation will differ.
	err := r.db.Model(&usermodel.UserInfo{}).Where("user_id = ?", userID).Update("info_id", infoID).Error
	if err != nil {
		return err
	}
	return nil
}
