package taskmanager

import "time"

func (tm *TasksManager) UpdateTask(id string, newDescription string) Task {
	var updatedTask Task

	for i := 0; i < len(tm.Tasks); i += 1 {
		if tm.Tasks[i].Id == id {
			updatedTask = tm.Tasks[i]
			updatedTask.Description = newDescription
			updatedTask.UpdatedAt = time.Now()

			return updatedTask
		}
	}

	return Task{}
}
