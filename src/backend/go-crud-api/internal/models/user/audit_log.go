package user_model

import (
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// AuditLog records actions performed within the system for auditing purposes.
// It logs the action taken, the entity affected, and the user who performed the action.
type AuditLog struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"` // Primary key with auto-increment
	Action      string    `json:"action" gorm:"size:50;not null"`     // Action performed (e.g., CREATE, UPDATE, DELETE), not null
	Entity      string    `json:"entity" gorm:"size:100;not null"`    // Entity affected by the action, not null
	EntityID    uint      `json:"entity_id" gorm:"not null"`          // ID of the affected entity, not null
	Changes     string    `json:"changes,omitempty" gorm:"type:text"` // Description of what was changed, text type
	PerformedBy uint      `json:"performed_by" gorm:"not null"`       // ID of the user who performed the action, not null
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`   // When the action was performed
}

// BeforeCreate is a GORM hook that will be triggered before a new record is inserted into the database
func (audit *AuditLog) BeforeCreate(tx *gorm.DB) (err error) {
	if os.Getenv("ENABLE_AUDIT_LOG") == "true" {
		audit.CreatedAt = time.Now()
		logEntityChanged(audit)
	}
	return nil
}

// logEntityChanged is a helper function to log the entity changes
func logEntityChanged(audit *AuditLog) {
	log.Println("Entity Changed:", audit.Entity, "ID:", audit.EntityID, "Changes:", audit.Changes)
}
