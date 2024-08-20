package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"` // Primary key with auto-increment.
	IsActive bool   `json:"is_active" gorm:"default:true"`       // Sets default value to true, indicating the user is active by default.
	Name     string `json:"name" gorm:"size:100"`                // Limits the name field to 100 characters for efficient database usage.
	Email    string `json:"email" gorm:"uniqueIndex;size:100"`   // Ensures the email field is unique and limited to 100 characters.
}
