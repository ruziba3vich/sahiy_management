package task

import "context"

type TaskFilter struct {
	AssigneeID *int64
	ReviewerID *int64
	SectionID  *int64
	Status     *int
	Priority   *int
	ParentID   *int64
	Page       int
	PageSize   int
}

type TaskListResult struct {
	Tasks      []*Task
	TotalCount int64
}

type Repository interface {
	CreateTask(ctx context.Context, task *Task) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id int64) error
	GetTaskByID(ctx context.Context, id int64) (*Task, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetTasksWithFilter(ctx context.Context, filter *TaskFilter) (*TaskListResult, error)
	GetTasksForCalendar(ctx context.Context, sectionID, assigneeID *int64, startDate, endDate string, status *int) ([]*Task, error)
	GetEndDateForTask(ctx context.Context, taskID int64) (*int64, error)
}
