package useraction

import "context"

type Repository interface {
	CreateUserAction(ctx context.Context, action *UserAction) (*UserAction, error)
	UpdateUserAction(ctx context.Context, action *UserAction) (*UserAction, error)
	DeleteUserAction(ctx context.Context, id int64) error
	GetUserActionByID(ctx context.Context, id int64) (*UserAction, error)
	GetAllUserActions(ctx context.Context) ([]*UserAction, error)
}
