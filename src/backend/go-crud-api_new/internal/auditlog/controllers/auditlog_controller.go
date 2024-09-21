package auditlog_controllers

import (
	"github.com/gin-gonic/gin"
	services "go-crud-api/internal/auditlog/services"
	"net/http"
)

const (
	errMsgRetrieveAuditLogs = "Failed to retrieve audit logs"
	errMsgStoreAuditLog     = "Failed to store audit log"
)

type AuditLogController struct {
	auditLogService *services.AuditLogService
}

func NewAuditLogController(auditLogService *services.AuditLogService) *AuditLogController {
	return &AuditLogController{auditLogService: auditLogService}
}

// GetAuditLogs retrieves logs from Redis
func (controller *AuditLogController) GetAuditLogs(c *gin.Context) {
	logs, err := controller.auditLogService.GetAllLogs()
	if err != nil {
		handleError(c, http.StatusInternalServerError, errMsgRetrieveAuditLogs)
		return
	}
	c.JSON(http.StatusOK, logs)
}

// StoreAuditLog saves a new audit log to Redis
func (controller *AuditLogController) StoreAuditLog(c *gin.Context) {
	var log services.AuditLogEntry
	if err := c.ShouldBindJSON(&log); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := controller.auditLogService.StoreLog(log); err != nil {
		handleError(c, http.StatusInternalServerError, errMsgStoreAuditLog)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Log stored successfully"})
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
