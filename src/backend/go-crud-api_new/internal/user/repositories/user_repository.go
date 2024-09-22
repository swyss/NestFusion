package repositories

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/gorm"
)

// DB Type alias for gorm.DB
type DB = gorm.DB

// UserRepository UserRepo is a struct that interacts with the user database table.
type UserRepository struct {
	db *DB
}

// NewUserRepository NewUserRepo creates a new instance of UserRepo.
func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database.
func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

// GetUserByID retrieves a user from the database by ID.
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.findUserByID(id, &user)
	return &user, err
}

// GetUserByEmail retrieves a user from the database by email.
// This is the method needed for the login process.
func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := repo.db.Find(&users)
	return users, result.Error
}

// Extracted function to find user by ID
func (repo *UserRepository) findUserByID(id uint, user *models.User) error {
	return repo.db.First(user, id).Error
}
