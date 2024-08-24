
package models
import "time"

type Task struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"` // Primary key with auto-increment.
	Name     string `json:"name" gorm:"size:100"`                // Limits the name field to 100 characters for efficient database usage.
  Created time.Time `json:"created" gorm:"autoCreateTime"` // autoCreateTime sets the current timestamp when the record is created
	Due     time.Time `json:"due"`
	IsFinished bool   `json:"isFinished" gorm:"default:false"`       // Sets default value to true, indicating the user is active by default.
}

