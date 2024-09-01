package repos

import (
	"errors"
	"go-crud-api/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAllTasks() []models.Task {
	var task []models.Task
	result := r.DB.Find(&task)
	if result.Error != nil {
		panic(result.Error)
	}
	return task
}

func (r *TaskRepository) CreateTask(task *models.Task) ([]models.Task, error) {
	result := r.DB.Create(&task)
	if result.Error != nil {
		return r.GetAllTasks(), result.Error
	}
	return r.GetAllTasks(), nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) ([]models.Task, error) {
	var DBtask = models.Task{ID: task.ID}
	result := r.DB.First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = task.IsFinished
	DBtask.Name = task.Name
	DBtask.Due = task.Due

	savedTask := r.DB.Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return r.GetAllTasks(), nil
}

func (r *TaskRepository) MarkTaskAsDone(id uint) ([]models.Task, error) {
	var DBtask = models.Task{ID: id}
	result := r.DB.First(&DBtask)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	DBtask.IsFinished = true
	savedTask := r.DB.Save(&DBtask)
	if savedTask.Error != nil {
		return nil, errors.New("Failed to update task!")
	}

	return r.GetAllTasks(), nil
}

func (r *TaskRepository) DeleteTask(id uint) ([]models.Task, error) {
	var task = models.Task{ID: id}
	result := r.DB.Delete(&task)
	if result.RowsAffected == 0 {
		return r.GetAllTasks(), errors.New("Failed to delete task!")
	}

	return r.GetAllTasks(), nil
}
