package controllers

import (
	"github.com/gin-gonic/gin"
	user_models "go-crud-api/internal/user/models"
	"go-crud-api/internal/user/services"
	"net/http"
)

// UserInfoController handles requests related to user information.
type UserInfoController struct {
	UserInfoService services.UserInfoService
}

// GetAllUserInfo fetches all user info records.
func (c *UserInfoController) GetAllUserInfo(ctx *gin.Context) {
	userInfo, err := c.UserInfoService.GetAllUserInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user info."})
		return
	}
	ctx.JSON(http.StatusOK, userInfo)
}

// CreateUserInfo creates a new user info record.
func (c *UserInfoController) CreateUserInfo(ctx *gin.Context) {
	var userInfo user_models.UserInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserInfoService.CreateUserInfo(&userInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user info."})
		return
	}
	ctx.JSON(http.StatusOK, userInfo)
}
