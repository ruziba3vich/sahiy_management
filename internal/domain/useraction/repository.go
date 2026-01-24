package useraction

import "context"

type Repository interface {
	CreateUserAction(ctx context.Context, action *UserAction) (*UserAction, error)
	UpdateUserAction(ctx context.Context, action *UserAction) (*UserAction, error)
	DeleteUserAction(ctx context.Context, id int64) error
	GetUserActionByID(ctx context.Context, id int64) (*UserAction, error)
	GetAllUserActions(ctx context.Context) ([]*UserAction, error)
	GetLastActiveActionByUserID(ctx context.Context, userID int64) (*UserAction, error)
	GetAttendance(ctx context.Context, branchID int64, fromDate, toDate int64) ([]*AttendanceRecord, error)
}

type AttendanceRecord struct {
	UserAction
	UserFullName string
	UserPhone    string
}
