package models

type UserRole struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}
