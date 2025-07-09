package taskmanager

import (
	"time"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TasksManager struct {
	LastUsedId int    `json:"lastUsedId"`
	Tasks      []Task `json:"tasks"`
}
