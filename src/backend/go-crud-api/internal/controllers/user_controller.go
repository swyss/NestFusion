package controllers

import (
	"encoding/json"
	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserController is responsible for handling user-related HTTP requests.
type UserController struct {
	Service *services.UserService
}

// NewUserController creates a new instance of UserController with the provided UserService.
func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get details of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		// If there is an error, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the list of users and a 200 status code.
	c.handleJSONResponse(w, users, http.StatusOK)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get details of a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := c.Service.GetUserByID(id)
	if err != nil {
		// If the user is not found, respond with a 404 status code.
		c.handleError(w, err, http.StatusNotFound)
		return
	}
	// Respond with the user details and a 200 status code.
	c.handleJSONResponse(w, user, http.StatusOK)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Create user"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If decoding the request body fails, do not proceed.
		return
	}
	err = c.Service.CreateUser(&user)
	if err != nil {
		// If there is an error while creating the user, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the created user and a 201 status code.
	c.handleJSONResponse(w, user, http.StatusCreated)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user by ID with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If decoding the request body fails, do not proceed.
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user.ID = uint(id)

	err = c.Service.UpdateUser(&user)
	if err != nil {
		// If there is an error while updating the user, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the updated user and a 200 status code.
	c.handleJSONResponse(w, user, http.StatusOK)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.Service.DeleteUser(id)
	if err != nil {
		// If the user is not found, respond with a 404 status code.
		c.handleError(w, err, http.StatusNotFound)
		return
	}
	// Respond with a success message and a 200 status code.
	c.handleJSONResponse(w, map[string]string{"message": "User deleted"}, http.StatusOK)
}

// handleError handles errors by sending appropriate HTTP responses.
func (c *UserController) handleError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

// handleJSONResponse encodes and sends a JSON response with a given status code.
func (c *UserController) handleJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	jsonEncoder := json.NewEncoder(w)
	if err := jsonEncoder.Encode(data); err != nil {
		// If encoding the response fails, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
	}
}
