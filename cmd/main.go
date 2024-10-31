package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aminulislamabid/task-tracker-cli/internal/constants"
	"github.com/aminulislamabid/task-tracker-cli/internal/storage"
	"github.com/aminulislamabid/task-tracker-cli/internal/task"
	"github.com/aminulislamabid/task-tracker-cli/internal/utils"
)

func main() {
	err := storage.LoadTasks(constants.Filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: <program-run-cmd> <command> [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: [program-run-cmd] add <description>")
			return
		}
		task.AddTask(os.Args[2], constants.Filename)
		fmt.Println("Task added successfully")

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: [program-run-cmd] update <id> <description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		err = task.UpdateTask(id, os.Args[3], constants.Filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task updated successfully")
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: [program-run-cmd] delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		err = task.DeleteTask(id, constants.Filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task deleted successfully")
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: [program-run-cmd] mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		err = task.MarkInProgress(id, constants.Filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task marked as in progress")
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: [program-run-cmd] mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		err = task.MarkDone(id, constants.Filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task marked as done")
		}

	case "list":
		if len(os.Args) == 3 {
			tasks := task.ListTasksByStatus(utils.ConvertTaskStatus(os.Args[2]))
			for _, t := range tasks {
				fmt.Printf("\nID: %d\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
					t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
			}
		} else {
			tasks := task.ListTasks()
			for _, t := range tasks {
				fmt.Printf("\nID: %d\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
					t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
			}
		}

	default:
		fmt.Println("Unknown command", command)
	}
}
