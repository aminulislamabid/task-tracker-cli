package utils

import (
	"os"
	"testing"

	"github.com/aminulislamabid/task-tracker-cli/internal/constants"
	"github.com/aminulislamabid/task-tracker-cli/internal/models"
)

// Setup for testing task features
func Setup() {
	models.Tasks = []models.Task{}
	os.Remove(constants.TestFilename)
}

// Create a temporary file for testing
func CreateTempFile(t *testing.T) *os.File {
	file, err := os.CreateTemp("./", "tasks_test_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return file
}

func FindTaskById(id int) *models.Task {
	for i := range models.Tasks {
		if models.Tasks[i].ID == id {
			return &models.Tasks[i]
		}
	}

	return nil
}

func ConvertTaskStatus(status string) models.TaskStatus {
	switch status {
	case "todo":
		return models.Todo
	case "in-progress":
		return models.InProgress
	case "done":
		return models.Done
	default:
		return models.Todo
	}
}
