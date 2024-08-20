package repos

import (
	"go-crud-api/internal/models"
	"gorm.io/gorm"
)

// UserRepository handles the CRUD operations for the User model in the database.
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository with the provided database connection.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	// GORM simplifies querying all records.
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a single user by their ID.
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	// Use GORM to find the user by primary key.
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(u *models.User) error {
	// GORM automatically handles the insert and ID generation.
	if err := r.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(u *models.User) error {
	// GORM simplifies the update process.
	if err := r.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database by their ID.
func (r *UserRepository) DeleteUser(id uint) error {
	// GORM provides a method for deleting by primary key.
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
