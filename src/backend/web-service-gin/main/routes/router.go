package routes

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Register routes
    RegisterTaskRoutes(r)

    return r
}
