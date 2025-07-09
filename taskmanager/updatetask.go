package taskmanager

import "time"

func (tm *TasksManager) UpdateTask(id string, newDescription string) Task {
	for i := 0; i < len(tm.Tasks); i += 1 {
		if tm.Tasks[i].Id == id {
			tm.Tasks[i].Description = newDescription
			tm.Tasks[i].UpdatedAt = time.Now()

			return tm.Tasks[i]
		}
	}

	return Task{}
}
