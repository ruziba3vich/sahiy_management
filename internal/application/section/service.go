package section

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/section"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, departmentID int64, name string) (*domain.Section, error) {
	sec := domain.NewSection(departmentID, name)
	return s.repo.CreateSection(ctx, sec)
}

func (s *Service) Update(ctx context.Context, id, departmentID int64, name string) (*domain.Section, error) {
	sec, err := s.repo.GetSectionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	sec.Update(departmentID, name)
	return s.repo.UpdateSection(ctx, sec)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteSection(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Section, error) {
	return s.repo.GetSectionByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Section, error) {
	return s.repo.GetAllSections(ctx)
}
