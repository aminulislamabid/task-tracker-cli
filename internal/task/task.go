package task

import (
	"errors"
	"time"

	"github.com/aminulislamabid/task-tracker-cli/internal/models"
	"github.com/aminulislamabid/task-tracker-cli/internal/storage"
	"github.com/aminulislamabid/task-tracker-cli/internal/utils"
)

func AddTask(description string, filename string) {
	task := models.Task{
		ID:          len(models.Tasks) + 1,
		Description: description,
		Status:      models.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	models.Tasks = append(models.Tasks, task)
	storage.SaveTasks(filename)
}

func UpdateTask(id int, description string, filename string) error {
	task := utils.FindTaskById(id)
	if task == nil {
		return errors.New("task not found")
	}

	task.Description = description
	task.UpdatedAt = time.Now()
	return storage.SaveTasks(filename)
}

func DeleteTask(id int, filename string) error {
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return storage.SaveTasks(filename)
		}
	}
	return errors.New("task not found")
}

func MarkInProgress(id int, filename string) error {
	task := utils.FindTaskById(id)
	if task == nil {
		return errors.New("task not found")
	}
	task.Status = models.InProgress
	task.UpdatedAt = time.Now()
	return storage.SaveTasks(filename)
}

func MarkDone(id int, filename string) error {
	task := utils.FindTaskById(id)
	if task == nil {
		return errors.New("task not found")
	}
	task.Status = models.Done
	task.UpdatedAt = time.Now()
	return storage.SaveTasks(filename)
}

func ListTasks() []models.Task {
	return models.Tasks
}

func ListTasksByStatus(status models.TaskStatus) []models.Task {
	var tasks []models.Task
	for _, task := range models.Tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
