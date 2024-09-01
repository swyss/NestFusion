package user_model

import (
	"gorm.io/gorm"
	"os"
)

// AccessLevel represents different levels of access within the system.
// This model is linked to roles, defining the permissions associated with each role.
// swagger:model
type AccessLevel struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"`                  // Primary key with auto increment.
	IsActive bool   `json:"is_active" gorm:"default:false"`                       // Indicates if the access level is active.
	Name     string `json:"name" gorm:"size:100;unique"`                          // Unique name of the access level.
	Roles    []Role `json:"roles,omitempty" gorm:"many2many:role_access_levels;"` // Associated roles.
}

// AfterCreate hook to log creation of access level
func (a *AccessLevel) AfterCreate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "CREATE",
			Entity:      "AccessLevel",
			EntityID:    a.ID,
			Changes:     "Created new access level",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterUpdate hook to log updates to access level
func (a *AccessLevel) AfterUpdate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "UPDATE",
			Entity:      "AccessLevel",
			EntityID:    a.ID,
			Changes:     "Updated access level details", // This could be enhanced to track specific field changes.
			PerformedBy: 0,                              // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete hook to log deletion of access level
func (a *AccessLevel) AfterDelete(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "DELETE",
			Entity:      "AccessLevel",
			EntityID:    a.ID,
			Changes:     "Deleted access level",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}
