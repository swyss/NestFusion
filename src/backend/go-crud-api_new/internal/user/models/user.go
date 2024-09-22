package user_models

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

// UserInfo represents additional user information, with a foreign key referencing the User.
type UserInfo struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Email  string `gorm:"type:varchar(100);unique;not null" json:"email"`
	User   User   `gorm:"foreignKey:UserID"`
}

// UserRole represents the role assigned to a user.
type UserRole struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"type:varchar(100);not null" json:"role_name"`
}

// LoginCredentials represents login information, not stored in the database.
type LoginCredentials struct {
	UserName string `gorm:"type:varchar(100);not null" json:"username"`
	Password string `json:"password"`
}
