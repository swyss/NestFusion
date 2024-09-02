package task_services

import (
	models "go-crud-api/internal/tasks/models"
	repositories "go-crud-api/internal/tasks/repositories"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (service *TaskService) RegisterTask(task *models.Task) error {
	return service.repo.CreateTask(task)
}

func (service *TaskService) GetTask(id uint) (*models.Task, error) {
	return service.repo.GetTaskByID(id)
}
