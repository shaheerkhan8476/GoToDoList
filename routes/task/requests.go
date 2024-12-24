package task

import (
	"github.com/google/uuid"
)

type Task struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Id          uuid.UUID `json:"id"`
}
