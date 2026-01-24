package task

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ruziba3vich/sahiy_management/internal/domain/section"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
	"github.com/ruziba3vich/sahiy_management/internal/domain/user"
)

type Service struct {
	repo            domain.Repository
	userRepo        user.Repository
	taskStatusRepo  taskstatus.Repository
	sectionRepo     section.Repository
	taskHistoryRepo taskhistory.Repository
}

func NewService(repo domain.Repository, userRepo user.Repository, taskStatusRepo taskstatus.Repository, sectionRepo section.Repository, taskHistoryRepo taskhistory.Repository) *Service {
	return &Service{
		repo:            repo,
		userRepo:        userRepo,
		taskStatusRepo:  taskStatusRepo,
		sectionRepo:     sectionRepo,
		taskHistoryRepo: taskHistoryRepo,
	}
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
	createdTask, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	// Create initial history
	// Note: We need a user_id for the history. Since this is likely called from a handler,
	// maybe the user_id should be passed. For now, we'll use 0 or handle it if available.
	// Looking at existing code, Create doesn't take currentUserID.
	// But history needs it. Let's check how it's handled in other places.
	// If we can't get it, we might need to update the signature or use a system ID.

	return createdTask, nil
}

func (s *Service) Update(ctx context.Context, id int64, parentID *int64, sectionID int64, assigneeID, reviewerID *int64, title, description string, priority int, deadline *int64, status int, currentUserID int64) (*domain.Task, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	oldStatus := task.Status

	// Check if status is being changed to Finished type
	if task.Status != status {
		newStatus, err := s.taskStatusRepo.GetTaskStatusByID(ctx, int64(status))
		if err == nil && newStatus.StatusType == taskstatus.Finished {
			// Get current user to check role and department
			currentUser, err := s.userRepo.GetUserByID(ctx, currentUserID)
			if err != nil {
				return nil, err
			}

			// Role must be 50 or 99 (Manager)
			if currentUser.Role != 50 && currentUser.Role != 99 {
				return nil, errors.New("only managers can set task to finished status")
			}

			// Must be in the same department
			// Fetch task's section to find out which department it belongs to
			taskSection, err := s.sectionRepo.GetSectionByID(ctx, task.SectionID)
			if err != nil {
				return nil, err
			}

			if taskSection.DepartmentID != currentUser.DepartmentID {
				return nil, errors.New("managers can only finish tasks in their own department")
			}
		}
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
	updatedTask, err := s.repo.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	// Update history if status changed
	if oldStatus != status {
		now := time.Now().Unix()
		// 1. Finish the old history record
		histories, err := s.taskHistoryRepo.GetTaskHistoriesWithFilter(ctx, &taskhistory.TaskHistoryFilter{
			TaskID: &id,
			Status: &oldStatus,
		})
		if err == nil && len(histories.Histories) > 0 {
			// Find the active one (where finished_at is null)
			for _, h := range histories.Histories {
				if !h.FinishedAt.Valid {
					h.FinishedAt = pgtype.Int8{Int64: now, Valid: true}
					s.taskHistoryRepo.UpdateTaskHistory(ctx, h)
					break
				}
			}
		}

		// 2. Create new history record
		newHistory := taskhistory.NewTaskHistory(id, currentUserID, status)
		newHistory.StartedAt = now
		s.taskHistoryRepo.CreateTaskHistory(ctx, newHistory)
	}

	return updatedTask, nil
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

func (s *Service) GetTasksForCalendar(ctx context.Context, sectionID, assigneeID *int64, startDate, endDate string, status *int) ([]*domain.Task, map[int64]*taskstatus.TaskStatus, map[int64]*section.Section, error) {
	tasks, err := s.repo.GetTasksForCalendar(ctx, sectionID, assigneeID, startDate, endDate, status)
	if err != nil {
		return nil, nil, nil, err
	}

	statusMap := make(map[int64]*taskstatus.TaskStatus)
	sectionMap := make(map[int64]*section.Section)

	for _, t := range tasks {
		if _, exists := statusMap[int64(t.Status)]; !exists {
			st, err := s.taskStatusRepo.GetTaskStatusByID(ctx, int64(t.Status))
			if err == nil {
				statusMap[int64(t.Status)] = st
			}
		}
		if _, exists := sectionMap[t.SectionID]; !exists {
			sec, err := s.sectionRepo.GetSectionByID(ctx, t.SectionID)
			if err == nil {
				sectionMap[t.SectionID] = sec
			}
		}
	}

	return tasks, statusMap, sectionMap, nil
}

func (s *Service) GetCalendarEventByID(ctx context.Context, id int64) (*domain.Task, *taskstatus.TaskStatus, *section.Section, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	endDate, _ := s.repo.GetEndDateForTask(ctx, id)
	task.EndDate = endDate

	st, _ := s.taskStatusRepo.GetTaskStatusByID(ctx, int64(task.Status))
	sec, _ := s.sectionRepo.GetSectionByID(ctx, task.SectionID)

	return task, st, sec, nil
}
