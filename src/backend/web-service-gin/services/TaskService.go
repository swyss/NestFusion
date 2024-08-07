package services

import (
	"web-service-gin/models"
	"web-service-gin/repositories"
)

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (service *TaskService) GetTasks() []models.Task {
	taskrepo := repositories.NewTaskRepository()

	return taskrepo.GetTasks()
}

func (service *TaskService) GetTaskById(id string) (*models.Task, error) {
	taskrepo := repositories.NewTaskRepository()
	return taskrepo.GetTaskById(id)
}
