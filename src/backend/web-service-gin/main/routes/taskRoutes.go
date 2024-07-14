package routes

import ("github.com/gin-gonic/gin" 
        "web-service-gin/controllers")

func RegisterTaskRoutes(r *gin.Engine){
  taskController := controllers.NewTaskController()
  taskRoutes := r.Group("/tasks")
  {
    taskRoutes.GET("", taskController.GetTasks)
    taskRoutes.GET("/:id", taskController.GetTaskById)
  }
}
