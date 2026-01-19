package privilege

import "time"

const (
	RoleUser       = 1
	RoleModerator  = 2
	RoleAdmin      = 3
	RoleSuperAdmin = 99
)

type Privilege struct {
	ID          int64
	Name        string
	Resource    string
	Action      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

func NewPrivilege(name, resource, action, description string) *Privilege {
	now := time.Now().Unix()
	return &Privilege{
		Name:        name,
		Resource:    resource,
		Action:      action,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (p *Privilege) Update(name, resource, action, description string) {
	p.Name = name
	p.Resource = resource
	p.Action = action
	p.Description = description
	p.UpdatedAt = time.Now().Unix()
}

type UserPrivilege struct {
	ID          int64
	UserID      int64
	PrivilegeID int64
	GrantedBy   *int64
	GrantedAt   int64
}

func NewUserPrivilege(userID, privilegeID int64, grantedBy *int64) *UserPrivilege {
	return &UserPrivilege{
		UserID:      userID,
		PrivilegeID: privilegeID,
		GrantedBy:   grantedBy,
		GrantedAt:   time.Now().Unix(),
	}
}
