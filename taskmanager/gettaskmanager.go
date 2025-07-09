package taskmanager

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func GetTaskManager() (*TasksManager, error) {
	jsonFile, err := os.Open(getFilePath())

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &TasksManager{}, nil
		}

		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var db TasksManager

	json.Unmarshal(byteValue, &db)

	return &db, nil
}
