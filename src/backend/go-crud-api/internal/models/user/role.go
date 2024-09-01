package user_model

import (
	"gorm.io/gorm"
	"os"
)

// Role represents a set of permissions assigned to users.
// Each role is associated with multiple access levels, defining the user's capabilities.
type Role struct {
	gorm.Model
	ID           uint          `json:"id" gorm:"primary_key;autoIncrement"`                          // Primary key with auto-increment.
	IsActive     bool          `json:"is_active" gorm:"default:false"`                               // Indicates if the role is active.
	Name         string        `json:"name" gorm:"size:100;unique"`                                  // Unique name of the role.
	AccessLevels []AccessLevel `json:"access_levels,omitempty" gorm:"many2many:role_access_levels;"` // Associated access levels.
	Description  string        `json:"description,omitempty" gorm:"size:255"`                        // Description of the role.
}

type AssignAccessLevelInput struct {
	AccessLevelID uint `json:"access_level_id" binding:"required"`
}

type LevelInput struct {
	RoleID        uint `json:"role_id" binding:"required" gorm:"not null"`  // Not null constraint for user ID.
	AccessLevelID uint `json:"level_id" binding:"required" gorm:"not null"` // Not null constraint for password ID.
}

// AfterCreate hook to log creation of role
func (r *Role) AfterCreate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "CREATE",
			Entity:      "Role",
			EntityID:    r.ID,
			Changes:     "Created new role",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterUpdate hook to log updates to role
func (r *Role) AfterUpdate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "UPDATE",
			Entity:      "Role",
			EntityID:    r.ID,
			Changes:     "Updated role details", // This could be enhanced to track specific field changes.
			PerformedBy: 0,                      // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete hook to log deletion of role
func (r *Role) AfterDelete(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		log := AuditLog{
			Action:      "DELETE",
			Entity:      "Role",
			EntityID:    r.ID,
			Changes:     "Deleted role",
			PerformedBy: 0, // Adjust to the actual user performing the action, if available.
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}
	return nil
}
