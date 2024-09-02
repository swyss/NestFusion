package task_repositories

import (
	models "go-crud-api/internal/tasks/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	return repo.db.Create(task).Error
}

func (repo *TaskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	err := repo.db.First(&task, id).Error
	return &task, err
}
