package user_controllers

import (
	"github.com/gin-gonic/gin"
	models "go-crud-api/internal/user/models"
	services "go-crud-api/internal/user/services"
	"net/http"
)

type UserInfoController struct {
	service *services.UserInfoService
}

func NewUserInfoController(service *services.UserInfoService) *UserInfoController {
	return &UserInfoController{service: service}
}

// CreateUserInfo handles the creation of user info.
func (controller *UserInfoController) CreateUserInfo(c *gin.Context) {
	var info models.UserInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.service.CreateUserInfo(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user info"})
		return
	}
	c.JSON(http.StatusCreated, info)
}
