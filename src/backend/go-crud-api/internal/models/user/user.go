package user_model

import (
	"gorm.io/gorm"
	"os"
)

// User represents a system user, containing authentication and role information.
type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`  // Primary key with auto-increment and not null constraint.
	Username string `json:"username" gorm:"uniqueIndex;size:50;not null"` // Unique index, size limited to 50 characters, and not null constraint.
	IsActive bool   `json:"is_active" gorm:"default:true;not null"`       // Default value set to true and not null constraint.
}

// AuthInput represents the input data required for authentication.
type AuthInput struct {
	UserID     uint `json:"user_id" binding:"required" gorm:"not null"`     // Not null constraint for user ID.
	PasswordID uint `json:"password_id" binding:"required" gorm:"not null"` // Not null constraint for password ID.
}

// RoleInput represents the input data for role assignment.
type RoleInput struct {
	UserID uint `json:"user_id" binding:"required" gorm:"not null"` // Not null constraint for user ID.
	RoleID uint `json:"role_id" binding:"required" gorm:"not null"` // Not null constraint for role ID.
}

// InfoInput represents the input data for additional user information.
type InfoInput struct {
	UserID uint `json:"user_id" binding:"required" gorm:"not null"` // Not null constraint for user ID.
	InfoID uint `json:"info_id" binding:"required" gorm:"not null"` // Not null constraint for info ID.
}

// AfterCreate hook to log creation.
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		auditlog := AuditLog{
			Action:      "CREATE",
			Entity:      "User",
			EntityID:    u.ID,
			Changes:     "Created new user",
			PerformedBy: u.ID, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&auditlog).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterUpdate hook to log updates.
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		auditlog := AuditLog{
			Action:      "UPDATE",
			Entity:      "User",
			EntityID:    u.ID,
			Changes:     "Updated user details", // This could be enhanced to track specific field changes.
			PerformedBy: u.ID,
		}
		if err := tx.Create(&auditlog).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete hook to log deletions.
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		auditlog := AuditLog{
			Action:      "DELETE",
			Entity:      "User",
			EntityID:    u.ID,
			Changes:     "Deleted user",
			PerformedBy: u.ID,
		}
		if err := tx.Create(&auditlog).Error; err != nil {
			return err
		}
	}
	return nil
}

// BeforeCreate hook to hash passwords or perform other pre-processing tasks.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Here you could add logic like password hashing or validations.
	return nil
}

// BeforeUpdate hook for pre-processing before updating the user.
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Here you could add logic like validations before updating.
	return nil
}
