package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "go-crud-api/internal/user/models"
	services "go-crud-api/internal/user/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := controller.service.RegisterUser(&user); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (controller *UserController) GetUser(c *gin.Context) {
	userID, err := parseUserID(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := controller.service.GetUserByID(userID)
	if err != nil {
		handleError(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := controller.service.RegisterUser(&user); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to register user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	token, err := controller.service.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		handleError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *UserController) GetAllUsers(c *gin.Context) {
	users, err := controller.service.GetAllUsers()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	c.JSON(http.StatusOK, users)
}

func handleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func parseUserID(idParam string) (uint, error) {
	const base = 10
	parsedID, err := strconv.ParseUint(idParam, base, 32)
	return uint(parsedID), err
}
