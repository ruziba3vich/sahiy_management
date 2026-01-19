package privilege

import "context"

type Repository interface {
	CreatePrivilege(ctx context.Context, privilege *Privilege) (*Privilege, error)
	UpdatePrivilege(ctx context.Context, privilege *Privilege) (*Privilege, error)
	DeletePrivilege(ctx context.Context, id int64) error
	GetPrivilegeByID(ctx context.Context, id int64) (*Privilege, error)
	GetAllPrivileges(ctx context.Context) ([]*Privilege, error)
	GetPrivilegeByResourceAction(ctx context.Context, resource, action string) (*Privilege, error)

	AssignPrivilegeToUser(ctx context.Context, up *UserPrivilege) (*UserPrivilege, error)
	RevokePrivilegeFromUser(ctx context.Context, userID, privilegeID int64) error
	GetUserPrivileges(ctx context.Context, userID int64) ([]*Privilege, error)
	HasPrivilege(ctx context.Context, userID int64, resource, action string) (bool, error)
}
