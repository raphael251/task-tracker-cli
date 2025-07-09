package taskmanager

import (
	"encoding/json"
	"os"
)

func (tm *TasksManager) Save() error {
	marshalledDB, err := json.Marshal(tm)

	if err != nil {
		return err
	}

	createdFile, err := os.Create(getFilePath())

	if err != nil {
		return err
	}

	createdFile.Write(marshalledDB)

	return nil
}
