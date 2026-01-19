package branch

import "context"

type Repository interface {
	CreateBranch(ctx context.Context, branch *Branch) (*Branch, error)
	UpdateBranch(ctx context.Context, branch *Branch) (*Branch, error)
	DeleteBranch(ctx context.Context, id int64) error
	GetBranchByID(ctx context.Context, id int64) (*Branch, error)
	GetAllBranches(ctx context.Context) ([]*Branch, error)
}
