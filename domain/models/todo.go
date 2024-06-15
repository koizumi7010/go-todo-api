package models

import "time"

type ToDo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    TaskStatus
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

func NewTodo(task string) *ToDo {
	return &ToDo{
		Task:   task,
		Status: Created,
	}
}

func NewUpdateTodo(id int, task string, status TaskStatus) *ToDo {
	return &ToDo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

type TaskStatus string

const (
	Created    = TaskStatus("created")
	Processing = TaskStatus("processing")
	Done       = TaskStatus("done")
)
