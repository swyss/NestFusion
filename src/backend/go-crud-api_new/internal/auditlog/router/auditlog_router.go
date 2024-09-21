package auditlog_router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	controllers "go-crud-api/internal/auditlog/controllers"
	repositories "go-crud-api/internal/auditlog/repositories"
	services "go-crud-api/internal/auditlog/services"
)

const auditLogRoutePath = "/audit-logs"

func initAuditLogComponents() (*controllers.AuditLogController, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	auditLogRepo := repositories.NewAuditLogRepository(redisClient)
	auditLogService := services.NewAuditLogService(auditLogRepo)
	return controllers.NewAuditLogController(auditLogService), nil
}

func RegisterAuditLogRoutes(r *gin.RouterGroup) {
	auditLogController, err := initAuditLogComponents()
	if err != nil {
		return
	}

	r.GET(auditLogRoutePath, auditLogController.GetAuditLogs)
	r.POST(auditLogRoutePath, auditLogController.StoreAuditLog)
}
