package auditlog_services

import (
	models "go-crud-api/internal/auditlog/models"
	repositories "go-crud-api/internal/auditlog/repositories"
	"time"
)

type AuditLogEntry struct {
	UserID     uint   `json:"user_id"`
	ActionType string `json:"action_type"`
}

// AuditLogService provides services related to audit logs
type AuditLogService struct {
	repository *repositories.AuditLogRepository
}

func NewAuditLogService(repository *repositories.AuditLogRepository) *AuditLogService {
	return &AuditLogService{repository: repository}
}

// StoreLog stores a new audit log
func (service *AuditLogService) StoreLog(entry AuditLogEntry) error {
	log := models.AuditLog{
		UserID:     entry.UserID,
		ActionType: entry.ActionType,
		CreatedAt:  time.Now(),
	}
	return service.repository.StoreLog(log)
}

// GetAllLogs retrieves all audit logs
func (service *AuditLogService) GetAllLogs() ([]models.AuditLog, error) {
	return service.repository.GetAll()
}
