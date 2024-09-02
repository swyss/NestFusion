package auditlog_controllers

import (
	"github.com/gin-gonic/gin"
	services "go-crud-api/internal/auditlog/services"
	"net/http"
)

const (
	errMsgRetrieveAuditLogs = "Failed to retrieve audit logs"
)

type AuditLogController struct {
	auditLogService *services.AuditLogService
}

func NewAuditLogController(auditLogService *services.AuditLogService) *AuditLogController {
	return &AuditLogController{auditLogService: auditLogService}
}

func (controller *AuditLogController) GetAuditLogs(c *gin.Context) {
	logs, err := controller.auditLogService.GetAllLogs()
	if err != nil {
		handleError(c, http.StatusInternalServerError, errMsgRetrieveAuditLogs)
		return
	}
	c.JSON(http.StatusOK, logs)
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
