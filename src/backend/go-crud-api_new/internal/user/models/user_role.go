package models

import "gorm.io/gorm"

// UserRole represents the role assigned to a user.
type UserRole struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"type:varchar(100);not null" json:"role_name"`
}
