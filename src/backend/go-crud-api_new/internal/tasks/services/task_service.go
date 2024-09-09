package task_services

import (
	models "go-crud-api/internal/tasks/models"
  "go-crud-api/internal/tasks/repositories"
)

type TaskService struct {}
var _taskRepo *task_repositories.TaskRepository

func NewTaskService() *TaskService {
  _taskRepo = task_repositories.NewTaskRepository()
	return &TaskService{}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return _taskRepo.GetAllTasks(), nil
}

func (s *TaskService) CreateTask(t *models.Task) ([]models.Task, error) {
	return _taskRepo.CreateTask(t)
}

func (s *TaskService) UpdateTask(t *models.Task) ([]models.Task, error) {
	return _taskRepo.UpdateTask(t)
}

func (s *TaskService) MarkTaskAsDone(id uint) ([]models.Task, error) {
	return _taskRepo.MarkTaskAsDone(id)
}

func (s *TaskService) DeleteTask(id uint) ([]models.Task, error) {
	return _taskRepo.DeleteTask(id)
}
