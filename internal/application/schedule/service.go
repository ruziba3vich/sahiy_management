package schedule

import (
	"context"
	"encoding/json"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/schedule"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, timezone int, weekDays json.RawMessage, status int) (*domain.Schedule, error) {
	schedule := domain.NewSchedule(name, timezone, weekDays, status)
	return s.repo.CreateSchedule(ctx, schedule)
}

func (s *Service) Update(ctx context.Context, id int64, name string, timezone int, weekDays json.RawMessage, status int) (*domain.Schedule, error) {
	schedule, err := s.repo.GetScheduleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	schedule.Update(name, timezone, weekDays, status)
	return s.repo.UpdateSchedule(ctx, schedule)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteSchedule(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Schedule, error) {
	return s.repo.GetScheduleByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Schedule, error) {
	return s.repo.GetAllSchedules(ctx)
}
