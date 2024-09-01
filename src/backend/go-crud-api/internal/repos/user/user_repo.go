package user_repository

import (
	usermodel "go-crud-api/internal/models/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *usermodel.User) error
	GetUserByID(id uint) (*usermodel.User, error)
	UpdateUser(user *usermodel.User) error
	DeleteUser(id uint) error
	GetUsers() ([]usermodel.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *usermodel.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*usermodel.User, error) {
	var user usermodel.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *usermodel.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&usermodel.User{}, id).Error
}

func (r *userRepository) GetUsers() ([]usermodel.User, error) {
	var users []usermodel.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
