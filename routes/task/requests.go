package task

import (
	"github.com/google/uuid"
)

type Task struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Id          uuid.UUID `json:"id"`
}

func NewTask(name string, description string) Task {
	return Task{
		Name:        name,
		Description: description,
		Id:          uuid.New(),
	}
}
