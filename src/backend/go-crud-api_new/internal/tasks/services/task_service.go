package task_services

import (
	models "go-crud-api/internal/tasks/models"
  "go-crud-api/internal/tasks/repositories"
)

type TaskService struct {
	repository *task_repositories.TaskRepository
}

func NewTaskService(repo *task_repositories.TaskRepository) *TaskService {
	return &TaskService{repository: repo}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repository.GetAllTasks(), nil
}

func (s *TaskService) CreateTask(t *models.Task) ([]models.Task, error) {
	return s.repository.CreateTask(t)
}

func (s *TaskService) UpdateTask(t *models.Task) ([]models.Task, error) {
	return s.repository.UpdateTask(t)
}

func (s *TaskService) MarkTaskAsDone(id uint) ([]models.Task, error) {
	return s.repository.MarkTaskAsDone(id)
}

func (s *TaskService) DeleteTask(id uint) ([]models.Task, error) {
	return s.repository.DeleteTask(id)
}
