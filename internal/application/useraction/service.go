package useraction

import (
	"context"
	"errors"

	"math"
	"time"

	branchDomain "github.com/ruziba3vich/sahiy_management/internal/domain/branch"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/useraction"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
)

type Service struct {
	repo       domain.Repository
	branchRepo branchDomain.Repository
}

func NewService(repo domain.Repository, branchRepo branchDomain.Repository) *Service {
	return &Service{repo: repo, branchRepo: branchRepo}
}

func (s *Service) Create(ctx context.Context, userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) (*domain.UserAction, error) {
	action := domain.NewUserAction(userID, visitBranchID, leaveBranchID, comeStatus, outStatus, startedAt, finishedAt)
	return s.repo.CreateUserAction(ctx, action)
}

func (s *Service) Update(ctx context.Context, id int64, userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) (*domain.UserAction, error) {
	action, err := s.repo.GetUserActionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	action.Update(userID, visitBranchID, leaveBranchID, comeStatus, outStatus, startedAt, finishedAt)
	return s.repo.UpdateUserAction(ctx, action)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteUserAction(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.UserAction, error) {
	return s.repo.GetUserActionByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.UserAction, error) {
	return s.repo.GetAllUserActions(ctx)
}

func (s *Service) Accept(ctx context.Context, userID int64, lat, long float64) (*domain.UserAction, error) {
	branches, err := s.branchRepo.GetAllBranches(ctx)
	if err != nil {
		return nil, err
	}

	var nearBranchID int64 = -1
	for _, b := range branches {
		if s.calculateDistance(lat, long, b.Lat, b.Long) <= float64(b.Radius) {
			nearBranchID = b.ID
			break
		}
	}

	if nearBranchID == -1 {
		return nil, errors.New("user is not within any branch radius")
	}

	lastAction, err := s.repo.GetLastActiveActionByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, postgres.ErrUserActionNotFound) {
			now := time.Now().Unix()
			action := domain.NewUserAction(userID, nearBranchID, nil, 1, nil, now, nil)
			return s.repo.CreateUserAction(ctx, action)
		}
		return nil, err
	}

	now := time.Now().Unix()
	lastAction.FinishedAt = &now
	lastAction.LeaveBranchID = &nearBranchID
	status := 1
	lastAction.OutStatus = &status
	return s.repo.UpdateUserAction(ctx, lastAction)
}

func (s *Service) calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	phi1 := lat1 * math.Pi / 180
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func (s *Service) GetAttendance(ctx context.Context, branchID int64, fromDate, toDate int64) ([]*domain.AttendanceRecord, error) {
	return s.repo.GetAttendance(ctx, branchID, fromDate, toDate)
}

func (s *Service) GetStatuses() map[int]string {
	return domain.StatusLabels
}
