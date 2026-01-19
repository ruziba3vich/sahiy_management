package task

import "context"

type Repository interface {
	CreateTask(ctx context.Context, task *Task) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id int64) error
	GetTaskByID(ctx context.Context, id int64) (*Task, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
}
