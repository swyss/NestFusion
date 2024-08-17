package tests

import (
	"github.com/stretchr/testify/mock"
	"go-crud-api/internal/models"
)

// MockUserService implementiert das UserServiceInterface
type MockUserService struct {
	mock.Mock
}

// GetAllUsers simuliert die Methode zum Abrufen aller Benutzer
func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

// GetUserByID simuliert die Methode zum Abrufen eines Benutzers nach ID
func (m *MockUserService) GetUserByID(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

// CreateUser simuliert die Methode zum Erstellen eines neuen Benutzers
func (m *MockUserService) CreateUser(u *models.User) error {
	args := m.Called(u)
	return args.Error(0)
}

// UpdateUser simuliert die Methode zum Aktualisieren eines Benutzers
func (m *MockUserService) UpdateUser(u *models.User) error {
	args := m.Called(u)
	return args.Error(0)
}

// DeleteUser simuliert die Methode zum Löschen eines Benutzers
func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
