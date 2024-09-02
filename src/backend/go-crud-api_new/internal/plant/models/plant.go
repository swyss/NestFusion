package models

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserID      uint   `json:"user_id"`
}

// Constants for field names.
const (
	FieldName        = "Name"
	FieldDescription = "Description"
	FieldLocation    = "Location"
	FieldUserID      = "UserID"
)

// ValidateLocation validates the plant's location.
func ValidateLocation(location string) error {
	// Implement validation logic here
	return nil
}
