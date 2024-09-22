package user_controllers

import (
	"github.com/gin-gonic/gin"
	models "go-crud-api/internal/user/models"
	services "go-crud-api/internal/user/services"
	"net/http"
)

type UserRoleController struct {
	service *services.UserRoleService
}

func NewUserRoleController(service *services.UserRoleService) *UserRoleController {
	return &UserRoleController{service: service}
}

// CreateUserRole handles the creation of a new role.
func (controller *UserRoleController) CreateUserRole(c *gin.Context) {
	var role models.UserRole
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.service.CreateUserRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}
	c.JSON(http.StatusCreated, role)
}
