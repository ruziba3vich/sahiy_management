package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}
