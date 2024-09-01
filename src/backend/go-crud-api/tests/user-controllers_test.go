package tests

/*
import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	usercontroller "go-crud-api/internal/controllers/user"
	usermodel "go-crud-api/internal/models/user"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

// TestGetUsers tests the GetUsers method of the UserController
func TestGetUsers(t *testing.T) {
	mockService := new(MockUserService)
	controller := usercontroller.NewUserController(mockService)

	// Mock users with "is_active" field
	mockUsers := []usermodel.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", IsActive: false},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", IsActive: false},
	}
	mockService.On("GetAllUsers").Return(mockUsers, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUsers)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	expected := []map[string]interface{}{
		{"id": float64(1), "name": "John Doe", "email": "john@example.com", "is_active": false},
		{"id": float64(2), "name": "Jane Doe", "email": "jane@example.com", "is_active": false},
	}
	assert.Equal(t, expected, response)

	mockService.AssertExpectations(t)
}

// TestCreateUser tests the CreateUser method of the UserController
func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	newUser := &usermodel.User{Name: "Alice Smith", Email: "alice@example.com"}
	mockService.On("CreateUser", newUser).Return(nil)

	controller := usercontroller.NewUserController(mockService)

	userJSON, err := json.Marshal(newUser)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateUser)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertCalled(t, "CreateUser", newUser)
	mockService.AssertExpectations(t)
}

// TestGetUserByID tests the GetUserByID method of the UserController
func TestGetUserByID(t *testing.T) {
	mockService := new(MockUserService)
	controller := usercontroller.NewUserController(mockService)

	mockUser := &usermodel.User{ID: 1, Name: "John Doe", Email: "john@example.com", IsActive: false}
	mockService.On("GetUserByID", uint(1)).Return(mockUser, nil)

	req, err := http.NewRequest("GET", "/users/1", nil)
	require.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(1)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUser)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	expected := map[string]interface{}{
		"id":        float64(1),
		"name":      "John Doe",
		"email":     "john@example.com",
		"is_active": false,
	}
	assert.Equal(t, expected, response)

	mockService.AssertExpectations(t)
}

// TestDeleteUser tests the DeleteUser method of the UserController
func TestDeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := usercontroller.NewUserController(mockService)

	mockService.On("DeleteUser", uint(1)).Return(nil)

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	require.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(1)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteUser)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockService.AssertExpectations(t)
}
*/
