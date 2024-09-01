package user_controllers

import (
	"github.com/gin-gonic/gin"
	usermodel "go-crud-api/internal/models/user"
	service "go-crud-api/internal/services/user"
	"net/http"
)

// AuthController handles authentication-related HTTP requests.
type AuthController struct {
	authService service.AuthServiceInterface
}

// NewAuthController creates a new AuthController.
func NewAuthController(authService service.AuthServiceInterface) *AuthController {
	return &AuthController{authService}
}

// Authenticate handles the authentication of a user.
func (ctrl *AuthController) Authenticate(c *gin.Context) {
	var input usermodel.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the authentication logic in the service
	err := ctrl.authService.Authenticate(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

// RoleController handles role assignment-related HTTP requests.
type RoleController struct {
	roleService service.RoleServiceInterface
}

// NewRoleController creates a new RoleController.
func NewRoleController(roleService service.RoleServiceInterface) *RoleController {
	return &RoleController{roleService}
}

// AssignRole handles the assignment of a role to a user.
func (ctrl *RoleController) AssignRole(c *gin.Context) {
	var input usermodel.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the role assignment logic in the service
	err := ctrl.roleService.AssignRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}

// InfoController handles user information-related HTTP requests.
type InfoController struct {
	infoService service.InfoServiceInterface
}

// NewInfoController creates a new InfoController.
func NewInfoController(infoService service.InfoServiceInterface) *InfoController {
	return &InfoController{infoService}
}

// SetUserInfo handles the setting of additional user information.
func (ctrl *InfoController) SetUserInfo(c *gin.Context) {
	var input usermodel.InfoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the info setting logic in the service
	err := ctrl.infoService.SetUserInfo(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User info set successfully"})
}
