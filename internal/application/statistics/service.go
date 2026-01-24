package statistics

import (
	"context"
	"time"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/statistics"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDashboardStats(ctx context.Context) (*domain.DashboardStats, error) {
	now := time.Now()
	// Start of current month
	fromDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
	// End of today (or now)
	toDate := now.Unix()

	return s.repo.GetDashboardStats(ctx, fromDate, toDate)
}
