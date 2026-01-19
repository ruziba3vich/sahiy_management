package taskstatus

import "context"

type Repository interface {
	CreateTaskStatus(ctx context.Context, ts *TaskStatus) (*TaskStatus, error)
	UpdateTaskStatus(ctx context.Context, ts *TaskStatus) (*TaskStatus, error)
	DeleteTaskStatus(ctx context.Context, id int64) error
	GetTaskStatusByID(ctx context.Context, id int64) (*TaskStatus, error)
	GetAllTaskStatuses(ctx context.Context) ([]*TaskStatus, error)
}
