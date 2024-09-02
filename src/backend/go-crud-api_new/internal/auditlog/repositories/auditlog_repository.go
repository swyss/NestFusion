package auditlog_repositories

import (
	models "go-crud-api/internal/auditlog/models"
	"gorm.io/gorm"
)

// AuditLogRepository handles the data access for audit logs
type AuditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository creates a new instance of AuditLogRepository
func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

// GetAll retrieves all audit logs from the database
func (r *AuditLogRepository) GetAll() ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Find(&logs).Error
	return logs, err
}
