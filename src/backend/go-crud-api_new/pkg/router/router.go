package router

import (
	"github.com/gin-gonic/gin"
	auditRouter "go-crud-api/internal/auditlog/router"
	plantRouter "go-crud-api/internal/plant/router"
	taskRouter "go-crud-api/internal/tasks/router"
	userRouter "go-crud-api/internal/user/router"
)

func SetupRoutes(r *gin.Engine) {
	// Define your routes here
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is healthy",
		})
	})
	// Register User Routes
	userGroup := r.Group("/user")
	userRouter.RegisterUserRoutes(userGroup)

	// Register Task Routes
	taskGroup := r.Group("/tasks")
	taskRouter.RegisterTaskRoutes(taskGroup)

	// Register Audit Log Routes
	auditGroup := r.Group("/auditlog")
	auditRouter.RegisterAuditLogRoutes(auditGroup)

	// Register Plant Routes
	plantGroup := r.Group("/plant")
	plantRouter.RegisterPlantRoutes(plantGroup)
}
