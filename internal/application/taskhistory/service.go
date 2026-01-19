package taskhistory

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, taskID, userID int64, status int) (*domain.TaskHistory, error) {
	th := domain.NewTaskHistory(taskID, userID, status)
	return s.repo.CreateTaskHistory(ctx, th)
}

func (s *Service) Update(ctx context.Context, id, taskID, userID int64, status int, finishedAt *int64) (*domain.TaskHistory, error) {
	th, err := s.repo.GetTaskHistoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	th.Update(taskID, userID, status, finishedAt)
	return s.repo.UpdateTaskHistory(ctx, th)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteTaskHistory(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.TaskHistory, error) {
	return s.repo.GetTaskHistoryByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.TaskHistory, error) {
	return s.repo.GetAllTaskHistories(ctx)
}
