package user_model

import (
	"gorm.io/gorm"
	"os"
)

// UserInfo stores additional information about the user.
// This includes contact details and personal identifiers.
type UserInfo struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key;autoIncrement"`
	UserID    uint   `json:"user_id" gorm:"uniqueIndex"`        // Foreign key to the associated user.
	Email     string `json:"email" gorm:"uniqueIndex;size:100"` // Email for contact information, must be unique.
	FirstName string `json:"first_name" gorm:"size:100"`        // First name of the user.
	LastName  string `json:"last_name" gorm:"size:100"`         // Last name of the user.
	Address   string `json:"address,omitempty" gorm:"size:255"` // Address of the user.
	Phone     string `json:"phone,omitempty" gorm:"size:20"`    // Contact phone number.
}

// AfterCreate hook to log creation of user info
func (u *UserInfo) AfterCreate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "CREATE",
			Entity:      "UserInfo",
			EntityID:    u.UserID,
			Changes:     "Created new user info",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterUpdate hook to log updates to user info
func (u *UserInfo) AfterUpdate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "UPDATE",
			Entity:      "UserInfo",
			EntityID:    u.UserID,
			Changes:     "Updated user info details", // This could be enhanced to track specific field changes.
			PerformedBy: 0,                           // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete hook to log deletion of user info
func (u *UserInfo) AfterDelete(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "DELETE",
			Entity:      "UserInfo",
			EntityID:    u.UserID,
			Changes:     "Deleted user info",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}
