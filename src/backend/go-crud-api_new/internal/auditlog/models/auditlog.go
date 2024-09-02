package auditlog_models

import (
	"gorm.io/gorm"
	"time"
)

type AuditLog struct {
	gorm.Model
	ActionType string
	CreatedAt  time.Time
	UserID     uint
}
