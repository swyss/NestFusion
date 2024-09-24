package models

import (
	"gorm.io/gorm"
	"time"
)

// UserInfo represents additional user information, with a foreign key referencing the User.
type UserInfo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	User      User   `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
