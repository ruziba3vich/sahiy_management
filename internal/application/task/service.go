package task

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, parentID *int64, sectionID int64, assigneeID, reviewerID *int64, title, description string, priority int, deadline *int64, status int) (*domain.Task, error) {
	var parent pgtype.Int8
	if parentID != nil {
		parent = pgtype.Int8{Int64: *parentID, Valid: true}
	}

	var assignee pgtype.Int8
	if assigneeID != nil {
		assignee = pgtype.Int8{Int64: *assigneeID, Valid: true}
	}

	var reviewer pgtype.Int8
	if reviewerID != nil {
		reviewer = pgtype.Int8{Int64: *reviewerID, Valid: true}
	}

	var deadlineVal pgtype.Int8
	if deadline != nil {
		deadlineVal = pgtype.Int8{Int64: *deadline, Valid: true}
	}

	task := domain.NewTask(parent, sectionID, assignee, reviewer, title, description, priority, deadlineVal, status)
	return s.repo.CreateTask(ctx, task)
}

func (s *Service) Update(ctx context.Context, id int64, parentID *int64, sectionID int64, assigneeID, reviewerID *int64, title, description string, priority int, deadline *int64, status int) (*domain.Task, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var parent pgtype.Int8
	if parentID != nil {
		parent = pgtype.Int8{Int64: *parentID, Valid: true}
	}

	var assignee pgtype.Int8
	if assigneeID != nil {
		assignee = pgtype.Int8{Int64: *assigneeID, Valid: true}
	}

	var reviewer pgtype.Int8
	if reviewerID != nil {
		reviewer = pgtype.Int8{Int64: *reviewerID, Valid: true}
	}

	var deadlineVal pgtype.Int8
	if deadline != nil {
		deadlineVal = pgtype.Int8{Int64: *deadline, Valid: true}
	}

	task.Update(parent, sectionID, assignee, reviewer, title, description, priority, deadlineVal, status)
	return s.repo.UpdateTask(ctx, task)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteTask(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.GetAllTasks(ctx)
}

func (s *Service) GetAllWithFilter(ctx context.Context, filter *domain.TaskFilter) (*domain.TaskListResult, error) {
	return s.repo.GetTasksWithFilter(ctx, filter)
}
