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
	UserService services.UserServiceInterface
}

// NewUserController creates a new instance of UserController with the provided UserServiceInterface.
func NewUserController(service services.UserServiceInterface) *UserController {
	return &UserController{UserService: service}
}

// GetUsers @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		// If there is an error, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the list of users and a 200 status code.
	c.handleJSONResponse(w, users, http.StatusOK)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}
	user, err := c.UserService.GetUserByID(uint(id))
	if err != nil {
		// If the user is not found, respond with a 404 status code.
		c.handleError(w, err, http.StatusNotFound)
		return
	}
	// Respond with the user details and a 200 status code.
	c.handleJSONResponse(w, user, http.StatusOK)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If decoding the request body fails, respond with a 400 status code.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}
	err = c.UserService.CreateUser(&user)
	if err != nil {
		// If there is an error while creating the user, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the created user and a 201 status code.
	c.handleJSONResponse(w, user, http.StatusCreated)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If decoding the request body fails, respond with a 400 status code.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}
	user.ID = uint(id)

	err = c.UserService.UpdateUser(&user)
	if err != nil {
		// If there is an error while updating the user, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the updated user and a 200 status code.
	c.handleJSONResponse(w, user, http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = c.UserService.DeleteUser(uint(id))
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
