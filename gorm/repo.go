package gorm


import (
	"github.com/jinzhu/gorm"
	"github.com/vtomkiv/golang.api/api"
)

// TaskService represents a MySQL implementation of TaskService
type TaskRepository struct {
	DB *gorm.DB
}

// CreateTask persists a task and returns a assigned id.
func (r *TaskRepository) CreateTask(t *api.Task) (int64, error){

	// according to doc CreatedAt, UpdatedAt will be set automatically
	r.DB.Save(t)

	return t.Id, nil
}

// Task returns a task for a given id.
func (r *TaskRepository) GetTask(id int64) (*api.Task, error) {
	var task api.Task
	r.DB.First(&task, id)
	return &task, nil
}

// UpdateTask persists a task and returns the updated representation.
func (r *TaskRepository) UpdateTask(t *api.Task) (*api.Task, error){
	// according to doc UpdatedAt will be set automatically
	r.DB.Update(t)

	return t, nil
}

// DeleteTask marks a record as deleted
func (r *TaskRepository) DeleteTask(id int64) error{
	var task, _ = r.GetTask(id)

	// soft delete, according to doc UpdatedAt will be set automatically
	r.DB.Model(&task).Update("IsDeleted", true)

    return nil
}
