package department

import "context"

type Repository interface {
	CreateDepartment(ctx context.Context, dept *Department) (*Department, error)
	UpdateDepartment(ctx context.Context, dept *Department) (*Department, error)
	DeleteDepartment(ctx context.Context, id int64) error
	GetDepartmentByID(ctx context.Context, id int64) (*Department, error)
	GetAllDepartments(ctx context.Context) ([]*Department, error)

	AttachBranch(ctx context.Context, db *DepartmentBranch) (*DepartmentBranch, error)
	UpdateDepartmentBranchStatus(ctx context.Context, id int64, status int) error
	GetBranchesByDepartmentID(ctx context.Context, departmentID int64) ([]*DepartmentBranch, error)
}
