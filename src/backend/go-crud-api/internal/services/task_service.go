package services

import (
	"go-crud-api/internal/models"
	"go-crud-api/internal/repos"
)

type TaskServiceInterface interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(t *models.Task) ([]models.Task, error)
	UpdateTask(t *models.Task) ([]models.Task, error)
  MarkTaskAsDone(id uint) ([]models.Task, error)
  DeleteTask(id uint) ([]models.Task, error)
}
type TaskService struct {
	Repo *repos.TaskRepository
}

func NewTaskService(repo *repos.TaskRepository) *TaskService{
  return &TaskService{Repo: repo}
}

var _ TaskServiceInterface = &TaskService{}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repo.GetAllTasks(), nil // Assuming Repo has a corresponding method
}

func (s *TaskService) CreateTask(t *models.Task) ([]models.Task, error) {
	return s.Repo.CreateTask(t) // Assuming Repo has a corresponding method
}

func (s *TaskService) UpdateTask(t *models.Task) ([]models.Task, error) {
	return s.Repo.UpdateTask(t) // Assuming Repo has a corresponding method
}

func (s *TaskService) MarkTaskAsDone(id uint) ([]models.Task, error) {
	return s.Repo.MarkTaskAsDone(id) // Assuming Repo has a corresponding method
}

func (s *TaskService) DeleteTask(id uint) ([]models.Task, error) {
	return s.Repo.DeleteTask(id) // Assuming Repo has a corresponding method
}

