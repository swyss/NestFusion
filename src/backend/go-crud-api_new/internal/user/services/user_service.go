package user_services

import (
	"errors"
	models "go-crud-api/internal/user/models"
	repositories "go-crud-api/internal/user/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// UserService provides user-related services.
type UserService struct {
	repository *repositories.UserRepository
}

const (
	Pepper    = "my_secret_pepper" // Use environment variable for production
	JwtSecret = "supersecretkey"
)

// NewUserService creates a new instance of UserService.
func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

// GetUserByID retrieves a user by their ID
func (service *UserService) GetUserByID(userID uint) (*models.User, error) {
	user, err := service.repository.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// RegisterUser hashes the password and saves the user
func (service *UserService) RegisterUser(user *models.User) error {
	hashedPassword, salt, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.Salt = salt
	return service.repository.CreateUser(user)
}

// LoginUser verifies credentials and returns a JWT token
func (service *UserService) LoginUser(email string, password string) (string, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !verifyPassword(password, user.Password, user.Salt) {
		return "", errors.New("invalid credentials")
	}

	return createJwtToken(user.ID)
}

// hashPassword generates a bcrypt hash of the password with salt and pepper
func hashPassword(password string) (string, string, error) {
	salt := os.Getenv("SALT")
	if salt == "" {
		salt = "default_salt" // Fallback für Salt
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(salt+password+Pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashedPassword), salt, nil
}

// verifyPassword checks if the password is correct
func verifyPassword(password string, hashedPassword string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(salt+password+Pepper))
	return err == nil
}

// createJwtToken creates a new JWT token
func createJwtToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
