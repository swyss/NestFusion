package task_repositories

import (
	"errors"
	models "go-crud-api/internal/tasks/models"
	"go-crud-api/pkg/databases/postgres"
)

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (repo *TaskRepository) GetAllTasks() []models.Task {
	var tasks []models.Task
	result := postgres.GetDBClient().Find(&tasks)
	if result.Error != nil {
		panic(result.Error)
	}
	return tasks
}

func (repo *TaskRepository) CreateTask(task *models.Task) ([]models.Task, error) {
	result := postgres.GetDBClient().Create(&task)
	if result.Error != nil {
		return repo.GetAllTasks(), result.Error
	}
	return repo.GetAllTasks(), nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) ([]models.Task, error) {
	var DBtask = models.Task{ID: task.ID}
	result := postgres.GetDBClient().First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = task.IsFinished
	DBtask.Name = task.Name
	DBtask.Due = task.Due

	savedTask := postgres.GetDBClient().Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return r.GetAllTasks(), nil
}

func (r *TaskRepository) MarkTaskAsDone(id uint) ([]models.Task, error) {
	var DBtask = models.Task{ID: id}
	result := postgres.GetDBClient().First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = true
	savedTask := postgres.GetDBClient().Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return r.GetAllTasks(), nil
}

func (r *TaskRepository) DeleteTask(id uint) ([]models.Task, error) {
	var task = models.Task{ID: id}
	result := postgres.GetDBClient().Delete(&task)
	if result.RowsAffected == 0 {
		return r.GetAllTasks(), errors.New("Failed to delete task!")
	}

	return r.GetAllTasks(), nil
}
