package auditlog_services

import (
	"fmt"
	models "go-crud-api/internal/auditlog/models"
	repositories "go-crud-api/internal/auditlog/repositories"
)

// Error messages
const (
	errMsgRetrieveAllLogs = "Failed to retrieve all audit logs"
)

type AuditLogService struct {
	repository *repositories.AuditLogRepository
}

func NewAuditLogService(repository *repositories.AuditLogRepository) *AuditLogService {
	return &AuditLogService{repository: repository}
}

func (service *AuditLogService) GetAllLogs() ([]models.AuditLog, error) {
	logs, err := service.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsgRetrieveAllLogs, err)
	}
	return logs, nil
}
