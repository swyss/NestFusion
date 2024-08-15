package controllers

import (
	"encoding/json"
	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"net/http"
)

type SettingController struct {
	SettingService *services.SettingService
}

func NewSettingController(service *services.SettingService) *SettingController {
	return &SettingController{SettingService: service}
}

// GetAllSettings godoc
// @Summary Get all settings
// @Description Get all settings in the system
// @Tags settings
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Setting
// @Failure 500 {object} map[string]string
// @Router /settings [get]
func (c *SettingController) GetAllSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := c.SettingService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(settings)
	if err != nil {
		return
	}
}

// CreateSetting godoc
// @Summary Create a new setting
// @Description Create a new setting in the system
// @Tags settings
// @Accept  json
// @Produce  json
// @Param setting body models.Setting true "Create Setting"
// @Success 201 {object} models.Setting
// @Failure 400 {object} map[string]string
// @Router /settings [post]
func (c *SettingController) CreateSetting(w http.ResponseWriter, r *http.Request) {
	var setting models.Setting
	if err := json.NewDecoder(r.Body).Decode(&setting); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.SettingService.Create(&setting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(setting)
	if err != nil {
		return
	}
}
