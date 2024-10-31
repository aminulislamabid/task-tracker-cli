package task

import (
	"testing"

	"github.com/aminulislamabid/task-tracker-cli/internal/constants"
	"github.com/aminulislamabid/task-tracker-cli/internal/models"
	"github.com/aminulislamabid/task-tracker-cli/internal/utils"
)

func TestAddTask(t *testing.T) {
	utils.Setup()
	AddTask("Test Task", constants.TestFilename)
	if len(models.Tasks) != 1 {
		t.Fatal("Expected 1 task")
	}
	if models.Tasks[0].Description != "Test Task" {
		t.Fatalf("Expected task description to be 'Test Task', got '%s'", models.Tasks[0].Description)
	}
}

func TestUpdateTask(t *testing.T) {
	utils.Setup()
	AddTask("Test Task", constants.TestFilename)
	err := UpdateTask(1, "Updated Task", constants.TestFilename)
	if err != nil {
		t.Fatal(err)
	}
	if models.Tasks[0].Description != "Updated Task" {
		t.Fatalf("Expected task description to be 'Updated Task', got '%s'", models.Tasks[0].Description)
	}
}

func TestDeleteTask(t *testing.T) {
	utils.Setup()
	AddTask("Test Task", constants.TestFilename)
	err := DeleteTask(1, constants.TestFilename)
	if err != nil {
		t.Fatal(err)
	}
	if len(models.Tasks) != 0 {
		t.Fatal("Expected 0 tasks after deletion")
	}
}

func TestMarkInProgress(t *testing.T) {
	utils.Setup()
	AddTask("Test Task", constants.TestFilename)
	err := MarkInProgress(1, constants.TestFilename)
	if err != nil {
		t.Fatal(err)
	}
	if models.Tasks[0].Status != models.InProgress {
		t.Fatalf("Expected task status to be '%s', got '%s'", models.InProgress, models.Tasks[0].Status)
	}
}

func TestMarkDone(t *testing.T) {
	utils.Setup()
	AddTask("Test Task", constants.TestFilename)
	err := MarkDone(1, constants.TestFilename)
	if err != nil {
		t.Fatal(err)
	}
	if models.Tasks[0].Status != models.Done {
		t.Fatalf("Expected task status to be '%s', got '%s'", models.Done, models.Tasks[0].Status)
	}
}

func TestListTasks(t *testing.T) {
	utils.Setup()
	AddTask("Test Task 1", constants.TestFilename)
	AddTask("Test Task 2", constants.TestFilename)
	tasks := ListTasks()
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Description != "Test Task 1" {
		t.Fatalf("Expected first task description to be 'Test Task 1', got '%s'", tasks[0].Description)
	}
	if tasks[1].Description != "Test Task 2" {
		t.Fatalf("Expected second task description to be 'Test Task 2', got '%s'", tasks[1].Description)
	}
}

func TestListTasksByStatus(t *testing.T) {
	utils.Setup()
	AddTask("Test Task 1", constants.TestFilename)
	AddTask("Test Task 2", constants.TestFilename)
	AddTask("Test Task 3", constants.TestFilename)
	MarkInProgress(1, constants.TestFilename)
	MarkDone(2, constants.TestFilename)
	inProgressTasks := ListTasksByStatus(models.InProgress)
	doneTasks := ListTasksByStatus(models.Done)
	todoTasks := ListTasksByStatus(models.Todo)
	if len(inProgressTasks) != 1 {
		t.Fatalf("Expected 1 in-progress task, got %d", len(inProgressTasks))
	}
	if inProgressTasks[0].Status != models.InProgress {
		t.Fatalf("Expected in-progress task status to be '%s', got '%s'", models.InProgress, inProgressTasks[0].Status)
	}
	if len(doneTasks) != 1 {
		t.Fatalf("Expected 1 done task, got %d", len(doneTasks))
	}
	if doneTasks[0].Status != models.Done {
		t.Fatalf("Expected done task status to be '%s', got '%s'", models.Done, doneTasks[0].Status)
	}
	if len(todoTasks) != 1 {
		t.Fatalf("Expected 1 todo task, got %d", len(todoTasks))
	}
	if todoTasks[0].Status != models.Todo {
		t.Fatalf("Expected todo task status to be '%s', got '%s'", models.Todo, todoTasks[0].Status)
	}

}
