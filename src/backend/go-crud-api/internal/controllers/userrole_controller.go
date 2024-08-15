package controllers

import (
	"encoding/json"
	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"net/http"
)

type UserRoleController struct {
	UserRoleService *services.UserRoleService
}

func NewUserRoleController(service *services.UserRoleService) *UserRoleController {
	return &UserRoleController{UserRoleService: service}
}

// GetAllRoles godoc
// @Summary Get all user roles
// @Description Get all user roles in the system
// @Tags userroles
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserRole
// @Failure 500 {object} map[string]string
// @Router /roles [get]
func (c *UserRoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := c.UserRoleService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(roles)
	if err != nil {
		return
	}
}

// CreateRole godoc
// @Summary Create a new user role
// @Description Create a new user role in the system
// @Tags userroles
// @Accept  json
// @Produce  json
// @Param role body models.UserRole true "Create UserRole"
// @Success 201 {object} models.UserRole
// @Failure 400 {object} map[string]string
// @Router /roles [post]
func (c *UserRoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.UserRole
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.UserRoleService.Create(&role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(role)
	if err != nil {
		return
	}
}
