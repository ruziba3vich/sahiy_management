package privilege

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/privilege"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, resource, action, description string) (*domain.Privilege, error) {
	p := domain.NewPrivilege(name, resource, action, description)
	return s.repo.CreatePrivilege(ctx, p)
}

func (s *Service) Update(ctx context.Context, id int64, name, resource, action, description string) (*domain.Privilege, error) {
	p, err := s.repo.GetPrivilegeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	p.Update(name, resource, action, description)
	return s.repo.UpdatePrivilege(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeletePrivilege(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Privilege, error) {
	return s.repo.GetPrivilegeByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Privilege, error) {
	return s.repo.GetAllPrivileges(ctx)
}

func (s *Service) AssignToUser(ctx context.Context, userID, privilegeID int64, grantedBy *int64) (*domain.UserPrivilege, error) {
	up := domain.NewUserPrivilege(userID, privilegeID, grantedBy)
	return s.repo.AssignPrivilegeToUser(ctx, up)
}

func (s *Service) RevokeFromUser(ctx context.Context, userID, privilegeID int64) error {
	return s.repo.RevokePrivilegeFromUser(ctx, userID, privilegeID)
}

func (s *Service) GetUserPrivileges(ctx context.Context, userID int64) ([]*domain.Privilege, error) {
	return s.repo.GetUserPrivileges(ctx, userID)
}
