package repositories

import (
	"errors"
	"time"
	"web-service-gin/models"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository{
  return &TaskRepository{}
}

var tasks = []models.Task{
  {ID: "1", Name: "Task One", Created: time.Now(), Due: time.Now().Add(24 * time.Hour), Finished: false},
  {ID: "2", Name: "Task Two", Created: time.Now(), Due: time.Now().Add(48 * time.Hour), Finished: true},
	{ID: "3", Name: "Task Three", Created: time.Now(), Due: time.Now().Add(72 * time.Hour), Finished: false},
}

func (repo *TaskRepository) GetTasks() ([]models.Task)  {
  return tasks
}

func (repo *TaskRepository) GetTaskById(id string) (*models.Task, error){
  for _, task := range tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, errors.New("Task not found")
}
