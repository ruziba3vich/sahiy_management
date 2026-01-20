package taskhistory

import "context"

type TaskHistoryFilter struct {
	UserID   *int64
	TaskID   *int64
	Status   *int
	Page     int
	PageSize int
}

type TaskHistoryListResult struct {
	Histories  []*TaskHistory
	TotalCount int64
}

type Repository interface {
	CreateTaskHistory(ctx context.Context, th *TaskHistory) (*TaskHistory, error)
	UpdateTaskHistory(ctx context.Context, th *TaskHistory) (*TaskHistory, error)
	DeleteTaskHistory(ctx context.Context, id int64) error
	GetTaskHistoryByID(ctx context.Context, id int64) (*TaskHistory, error)
	GetAllTaskHistories(ctx context.Context) ([]*TaskHistory, error)
	GetTaskHistoriesWithFilter(ctx context.Context, filter *TaskHistoryFilter) (*TaskHistoryListResult, error)
}
