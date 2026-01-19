package useraction

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/useraction"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) (*domain.UserAction, error) {
	action := domain.NewUserAction(userID, visitBranchID, leaveBranchID, comeStatus, outStatus, startedAt, finishedAt)
	return s.repo.CreateUserAction(ctx, action)
}

func (s *Service) Update(ctx context.Context, id int64, userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) (*domain.UserAction, error) {
	action, err := s.repo.GetUserActionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	action.Update(userID, visitBranchID, leaveBranchID, comeStatus, outStatus, startedAt, finishedAt)
	return s.repo.UpdateUserAction(ctx, action)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteUserAction(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.UserAction, error) {
	return s.repo.GetUserActionByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.UserAction, error) {
	return s.repo.GetAllUserActions(ctx)
}
