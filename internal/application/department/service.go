package department

import (
	"context"
	"time"

	global "github.com/ruziba3vich/sahiy_management/internal/domain"
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

func (s *Service) AttachBranch(ctx context.Context, branchID, departmentID int64) (*domain.DepartmentBranch, error) {
	db := &domain.DepartmentBranch{
		BranchID:     branchID,
		DepartmentID: departmentID,
		Status:       global.StatusActive,
		CreatedAt:    time.Now().Unix(),
	}
	return s.repo.AttachBranch(ctx, db)
}

func (s *Service) UpdateDepartmentBranchStatus(ctx context.Context, id int64, status int) error {
	return s.repo.UpdateDepartmentBranchStatus(ctx, id, status)
}

func (s *Service) GetBranchesByDepartmentID(ctx context.Context, departmentID int64) ([]*domain.DepartmentBranch, error) {
	return s.repo.GetBranchesByDepartmentID(ctx, departmentID)
}
