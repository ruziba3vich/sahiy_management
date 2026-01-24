package statistics

import "context"

type Repository interface {
	GetDashboardStats(ctx context.Context, fromDate, toDate int64) (*DashboardStats, error)
}
