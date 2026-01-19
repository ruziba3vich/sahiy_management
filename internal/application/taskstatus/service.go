package taskstatus

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, statusType int) (*domain.TaskStatus, error) {
	ts := domain.NewTaskStatus(name, statusType)
	return s.repo.CreateTaskStatus(ctx, ts)
}

func (s *Service) Update(ctx context.Context, id int64, name string, statusType int) (*domain.TaskStatus, error) {
	ts, err := s.repo.GetTaskStatusByID(ctx, id)
	if err != nil {
		return nil, err
	}

	ts.Update(name, statusType)
	return s.repo.UpdateTaskStatus(ctx, ts)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteTaskStatus(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.TaskStatus, error) {
	return s.repo.GetTaskStatusByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.TaskStatus, error) {
	return s.repo.GetAllTaskStatuses(ctx)
}
