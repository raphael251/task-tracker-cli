package taskmanager

import (
	"fmt"
	"time"
)

func (tm *TasksManager) AddTask(description string) Task {
	if len(tm.Tasks) == 0 {
		firstId := 1

		task := Task{
			fmt.Sprint(firstId),
			description,
			StatusTodo,
			time.Now().UTC(),
			time.Now().UTC(),
		}

		tm.LastUsedId = firstId
		tm.Tasks = []Task{
			task,
		}

		return task
	}

	incrementedId := tm.LastUsedId + 1

	tm.LastUsedId = incrementedId

	task := Task{
		fmt.Sprint(incrementedId),
		description,
		StatusTodo,
		time.Now(),
		time.Now(),
	}

	tm.Tasks = append(tm.Tasks, task)

	return task
}
