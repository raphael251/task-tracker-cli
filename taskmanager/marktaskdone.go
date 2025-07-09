package taskmanager

func (tm *TasksManager) MarkTaskDone(id string) Task {
	for i := 0; i < len(tm.Tasks); i++ {
		if tm.Tasks[i].Id == id {
			tm.Tasks[i].Status = StatusDone

			return tm.Tasks[i]
		}
	}

	return Task{}
}
