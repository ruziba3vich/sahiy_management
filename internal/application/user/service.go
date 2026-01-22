package user

import (
	"context"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/user"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, departmentID, sectionID int64, scheduleID *int64, role int, phone, fullName, passwordHash string, tgChatID int64) (*domain.User, error) {
	user := domain.NewUser(departmentID, sectionID, scheduleID, role, phone, fullName, passwordHash, tgChatID)
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) Update(ctx context.Context, id, departmentID, sectionID int64, scheduleID *int64, role int, phone, fullName string, joinedAt, tgChatID int64) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Update(departmentID, sectionID, scheduleID, role, phone, fullName, joinedAt, tgChatID)
	return s.repo.UpdateUser(ctx, user)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAllUsers(ctx)
}
