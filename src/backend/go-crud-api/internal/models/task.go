
package models

type Task struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"` // Primary key with auto-increment.
	Name     string `json:"name" gorm:"size:100"`                // Limits the name field to 100 characters for efficient database usage.
	IsFinished bool   `json:"isFinished" gorm:"default:false"`       // Sets default value to true, indicating the user is active by default.
}
