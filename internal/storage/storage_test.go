package storage

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/aminulislamabid/task-tracker-cli/internal/models"
	"github.com/aminulislamabid/task-tracker-cli/internal/utils"
)

func TestLoadTasks(t *testing.T) {
	tempFile := utils.CreateTempFile(t)
	defer func() {
		tempFile.Close()
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Fatalf("Failed to delete temp file: %v", err)
		}
	}()

	expectedTasks := []models.Task{
		{ID: 1, Description: "Test Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Description: "Test Task 2", Status: "in-progress", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// Write test data to the temp file
	data, err := json.Marshal(expectedTasks)
	if err != nil {
		t.Fatalf("Failed to marshal tasks: %v", err)
	}

	err = os.WriteFile(tempFile.Name(), data, 0644)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	err = LoadTasks(tempFile.Name())
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}

	// Verify that the loaded tasks match the original tasks
	if len(expectedTasks) != len(models.Tasks) {
		t.Errorf("Expected %d tasks, got %d", len(expectedTasks), len(models.Tasks))
	}

	for i, task := range expectedTasks {
		if task.ID != models.Tasks[i].ID || task.Description != models.Tasks[i].Description || task.Status != models.Tasks[i].Status {
			t.Errorf("Loaded task does not match expected task. Expected task %+v, got %+v", task, models.Tasks[i])
		}
	}
}

func TestSaveTasks(t *testing.T) {
	tempFile := utils.CreateTempFile(t)
	defer func() {
		tempFile.Close()
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Fatalf("Failed to delete temp file: %v", err)
		}
	}()

	models.Tasks = []models.Task{
		{
			ID:          1,
			Description: "Test Task 1",
			Status:      models.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Description: "Test Task 2",
			Status:      models.InProgress,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	err := SaveTasks(tempFile.Name())
	if err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read %s file: %v", tempFile.Name(), err)
	}

	var savedTasks []models.Task
	err = json.Unmarshal(data, &savedTasks)
	if err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	if len(savedTasks) != len(models.Tasks) {
		t.Fatalf("Expected %d tasks, got %d", len(models.Tasks), len(savedTasks))
	}
}
