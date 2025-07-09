package taskmanager

import (
	"slices"
)

func (tm *TasksManager) DeleteTask(id string) Task {

	for i := 0; i < len(tm.Tasks); i += 1 {
		if tm.Tasks[i].Id == id {
			deletedTask := tm.Tasks[i]
			tm.Tasks = slices.Delete(tm.Tasks, i, i+1)

			return deletedTask
		}
	}

	return Task{}
}
