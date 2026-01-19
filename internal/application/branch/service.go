package branch

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/branch"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, lat, long float64, radius, status int) (*domain.Branch, error) {
	branch := domain.NewBranch(name, lat, long, radius, status)
	return s.repo.CreateBranch(ctx, branch)
}

func (s *Service) Update(ctx context.Context, id int64, name string, lat, long float64, radius, status int) (*domain.Branch, error) {
	branch, err := s.repo.GetBranchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	branch.Update(name, lat, long, radius, status)
	return s.repo.UpdateBranch(ctx, branch)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteBranch(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Branch, error) {
	return s.repo.GetBranchByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Branch, error) {
	return s.repo.GetAllBranches(ctx)
}
