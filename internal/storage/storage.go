package storage

import (
	"encoding/json"
	"os"

	"github.com/aminulislamabid/task-tracker-cli/internal/models"
)

func LoadTasks(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &models.Tasks)
}

func SaveTasks(filename string) error {
	data, err := json.MarshalIndent(models.Tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
