package models

import "gorm.io/gorm"

// Length constants for various fields
const (
	MaxNameLength     = 100
	MaxEmailLength    = 100
	MaxPasswordLength = 255
)

// User represents a user in the system with a one-to-one relationship to a role.
type User struct {
	gorm.Model
	ID       uint     `gorm:"primaryKey" json:"id"`
	UserName string   `gorm:"type:varchar(100);not null" json:"username"`
	Password string   `gorm:"type:varchar(255);not null" json:"-"`
	Salt     string   `gorm:"type:varchar(255);not null" json:"-"`
	RoleID   uint     `gorm:"not null" json:"role_id"`
	Role     UserRole `gorm:"foreignKey:RoleID" json:"role"`
}

// LoginCredentials represents login information, not stored in the database.
type LoginCredentials struct {
	UserName string `gorm:"type:varchar(100);not null" json:"username"`
	Password string `json:"password"`
}
