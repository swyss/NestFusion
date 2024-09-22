package task_repositories

import (
	"errors"
	models "go-crud-api/internal/tasks/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (repo *TaskRepository) GetAllTasks() []models.Task {
	var tasks []models.Task
	result := repo.DB.Find(&tasks)
	if result.Error != nil {
		panic(result.Error)
	}
	return tasks
}

func (repo *TaskRepository) CreateTask(task *models.Task) ([]models.Task, error) {
	result := repo.DB.Create(&task)
	if result.Error != nil {
		return repo.GetAllTasks(), result.Error
	}
	return repo.GetAllTasks(), nil
}

func (repo *TaskRepository) UpdateTask(task *models.Task) ([]models.Task, error) {
	var DBtask = models.Task{ID: task.ID}
	result := repo.DB.First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = task.IsFinished
	DBtask.Name = task.Name
	DBtask.Due = task.Due

	savedTask := repo.DB.Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return repo.GetAllTasks(), nil
}

func (repo *TaskRepository) MarkTaskAsDone(id uint) ([]models.Task, error) {
	var DBtask = models.Task{ID: id}
	result := repo.DB.First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = true
	savedTask := repo.DB.Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return repo.GetAllTasks(), nil
}

func (repo *TaskRepository) DeleteTask(id uint) ([]models.Task, error) {
	var task = models.Task{ID: id}
	result := repo.DB.Delete(&task)
	if result.RowsAffected == 0 {
		return repo.GetAllTasks(), errors.New("Failed to delete task!")
	}

	return repo.GetAllTasks(), nil
}
