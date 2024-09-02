package user_controllers

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
	user, err := controller.service.GetUser(userID)
	if err != nil {
		handleError(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func handleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func parseUserID(idParam string) (uint, error) {
	const base = 10
	parsedID, err := strconv.ParseUint(idParam, base, 32)
	return uint(parsedID), err
}
