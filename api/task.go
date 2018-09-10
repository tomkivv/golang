package api

import (
	_ "github.com/jinzhu/gorm"
	"time"
)


type Task struct {
	Id          int64   `gorm:"primary_key;AUTO_INCREMENT"`
	Title       string
	Description string
	Priority    int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CompletedAt *time.Time
	IsDeleted   bool
	IsCompleted bool
}

type TaskService interface {
	GetTask(id int64) (*Task, error)
	CreateTask(t *Task) (int64, error)
	UpdateTask(t *Task) (*Task, error)
	DeleteTask(id int64)  error
}
