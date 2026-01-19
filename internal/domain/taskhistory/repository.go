package taskhistory

import "context"

type Repository interface {
	CreateTaskHistory(ctx context.Context, th *TaskHistory) (*TaskHistory, error)
	UpdateTaskHistory(ctx context.Context, th *TaskHistory) (*TaskHistory, error)
	DeleteTaskHistory(ctx context.Context, id int64) error
	GetTaskHistoryByID(ctx context.Context, id int64) (*TaskHistory, error)
	GetAllTaskHistories(ctx context.Context) ([]*TaskHistory, error)
}
