package auditlog_router

import (
	"github.com/gin-gonic/gin"
	controllers "go-crud-api/internal/auditlog/controllers"
	repositories "go-crud-api/internal/auditlog/repositories"
	services "go-crud-api/internal/auditlog/services"
	database "go-crud-api/pkg/databases/postgres"
)

// Constants for route paths
const auditLogRoutePath = "/"

// Function to initialize components
func initAuditLogComponents() (*controllers.AuditLogController, error) {
	auditLogRepo := repositories.NewAuditLogRepository(database.PostgresDB)
	auditLogService := services.NewAuditLogService(auditLogRepo)
	return controllers.NewAuditLogController(auditLogService), nil
}

func RegisterAuditLogRoutes(r *gin.RouterGroup) {
	auditLogController, err := initAuditLogComponents()
	if err != nil {
		// Handle initialization error (optional)
		return
	}
	r.GET(auditLogRoutePath, auditLogController.GetAuditLogs)
	// Additional audit log routes can be added here
}
