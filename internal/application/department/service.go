package department

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/department"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, status int) (*domain.Department, error) {
	dept := domain.NewDepartment(name, status)
	return s.repo.CreateDepartment(ctx, dept)
}

func (s *Service) Update(ctx context.Context, id int64, name string, status int) (*domain.Department, error) {
	dept, err := s.repo.GetDepartmentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dept.Update(name, status)
	return s.repo.UpdateDepartment(ctx, dept)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteDepartment(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Department, error) {
	return s.repo.GetDepartmentByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Department, error) {
	return s.repo.GetAllDepartments(ctx)
}
