package taskmanager

func (tm *TasksManager) ListTasks(status ...string) []Task {
	if len(tm.Tasks) == 0 {
		return nil
	}

	if len(status) == 0 {
		return tm.Tasks
	}

	var tasks []Task

	for _, task := range tm.Tasks {
		if task.Status == status[0] {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
