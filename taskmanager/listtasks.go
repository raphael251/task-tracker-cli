package taskmanager

func (tm *TasksManager) ListTasks() []Task {
	if len(tm.Tasks) == 0 {
		return nil
	}

	return tm.Tasks
}
