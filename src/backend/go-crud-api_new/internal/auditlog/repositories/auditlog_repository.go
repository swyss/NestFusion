package auditlog_repositories

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	models "go-crud-api/internal/auditlog/models"
)

var ctx = context.Background()

// AuditLogRepository handles data access for Redis
type AuditLogRepository struct {
	redisClient *redis.Client
}

// NewAuditLogRepository creates a new instance of AuditLogRepository
func NewAuditLogRepository(redisClient *redis.Client) *AuditLogRepository {
	return &AuditLogRepository{redisClient: redisClient}
}

// StoreLog saves a new audit log in Redis
func (r *AuditLogRepository) StoreLog(log models.AuditLog) error {
	logData, err := json.Marshal(log)
	if err != nil {
		return err
	}
	return r.redisClient.LPush(ctx, "audit_logs", logData).Err()
}

// GetAll retrieves all audit logs from Redis
func (r *AuditLogRepository) GetAll() ([]models.AuditLog, error) {
	logs, err := r.redisClient.LRange(ctx, "audit_logs", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var auditLogs []models.AuditLog
	for _, log := range logs {
		var auditLog models.AuditLog
		if err := json.Unmarshal([]byte(log), &auditLog); err == nil {
			auditLogs = append(auditLogs, auditLog)
		}
	}
	return auditLogs, nil
}
