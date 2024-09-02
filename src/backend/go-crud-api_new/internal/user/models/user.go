package user_models

import "gorm.io/gorm"

// Length constants for various fields
const (
	MaxNameLength     = 100
	MaxEmailLength    = 100
	MaxPasswordLength = 255
)

// User represents a user in the system with JSON tags for API compatibility.
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}
