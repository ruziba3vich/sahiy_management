package schedule

import "context"

type Repository interface {
	CreateSchedule(ctx context.Context, schedule *Schedule) (*Schedule, error)
	UpdateSchedule(ctx context.Context, schedule *Schedule) (*Schedule, error)
	DeleteSchedule(ctx context.Context, id int64) error
	GetScheduleByID(ctx context.Context, id int64) (*Schedule, error)
	GetAllSchedules(ctx context.Context) ([]*Schedule, error)
}
