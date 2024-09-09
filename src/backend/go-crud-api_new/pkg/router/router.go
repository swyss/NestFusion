package router

import (
	"github.com/gin-gonic/gin"
	auditRouter "go-crud-api/internal/auditlog/router"
	plantRouter "go-crud-api/internal/plant/router"
	taskRouter "go-crud-api/internal/tasks/router"
	userRouter "go-crud-api/internal/user/router"
)

func SetupRoutes(r *gin.Engine) {
	// Register User Routes
	userGroup := r.Group("/user")
	userRouter.RegisterUserRoutes(userGroup)

	// Register Task Routes
	taskRouter.RegisterTaskRoutes(r)

	// Register Audit Log Routes
	auditGroup := r.Group("/auditlog")
	auditRouter.RegisterAuditLogRoutes(auditGroup)

	// Register Plant Routes
	plantGroup := r.Group("/plant")
	plantRouter.RegisterPlantRoutes(plantGroup)
}
