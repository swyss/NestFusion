package controllers

import (
	"github.com/gin-gonic/gin"
	models "go-crud-api/internal/user/models"
	"go-crud-api/internal/user/services"
	"net/http"
)

// UserRoleController handles requests related to user roles.
type UserRoleController struct {
	UserRoleService services.UserRoleService
}

// GetAllRoles fetches all roles.
func (c *UserRoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := c.UserRoleService.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve roles."})
		return
	}
	ctx.JSON(http.StatusOK, roles)
}

// CreateRole creates a new user role.
func (c *UserRoleController) CreateRole(ctx *gin.Context) {
	var role models.UserRole
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserRoleService.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create role."})
		return
	}
	ctx.JSON(http.StatusOK, role)
}
