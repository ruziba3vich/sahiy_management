package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
	userDomain "github.com/ruziba3vich/sahiy_management/internal/domain/user"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) domain.Repository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	query := `
		INSERT INTO tasks (parent_id, section_id, assignee_id, reviewer_id, title, description, priority, deadline, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		task.ParentID,
		task.SectionID,
		task.AssigneeID,
		task.ReviewerID,
		task.Title,
		task.Description,
		task.Priority,
		task.Deadline,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&task.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	query := `
		UPDATE tasks
		SET parent_id = $1, section_id = $2, assignee_id = $3, reviewer_id = $4, title = $5, description = $6, priority = $7, deadline = $8, status = $9, updated_at = $10
		WHERE id = $11
	`

	result, err := r.db.Exec(ctx, query,
		task.ParentID,
		task.SectionID,
		task.AssigneeID,
		task.ReviewerID,
		task.Title,
		task.Description,
		task.Priority,
		task.Deadline,
		task.Status,
		task.UpdatedAt,
		task.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id int64) (*domain.Task, error) {
	query := `
		SELECT 
			t.id, t.parent_id, t.section_id, t.assignee_id, t.reviewer_id, t.title, t.description, t.priority, t.deadline, t.status, t.created_at, t.updated_at,
			ua.id, ua.department_id, ua.section_id, ua.schedule_id, ua.role, ua.phone, ua.full_name, ua.joined_at, ua.created_at, ua.updated_at, ua.tg_chat_id,
			ur.id, ur.department_id, ur.section_id, ur.schedule_id, ur.role, ur.phone, ur.full_name, ur.joined_at, ur.created_at, ur.updated_at, ur.tg_chat_id
		FROM tasks t
		LEFT JOIN users ua ON t.assignee_id = ua.id
		LEFT JOIN users ur ON t.reviewer_id = ur.id
		WHERE t.id = $1
	`

	task := &domain.Task{}
	var aID, aDepID, aSecID, aRole, aJoinedAt, aCreatedAt, aUpdatedAt, aTgChatID pgtype.Int8
	var aSchedID pgtype.Int8
	var aPhone, aFullName pgtype.Text
	var rID, rDepID, rSecID, rRole, rJoinedAt, rCreatedAt, rUpdatedAt, rTgChatID pgtype.Int8
	var rSchedID pgtype.Int8
	var rPhone, rFullName pgtype.Text

	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.ParentID,
		&task.SectionID,
		&task.AssigneeID,
		&task.ReviewerID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
		&aID, &aDepID, &aSecID, &aSchedID, &aRole, &aPhone, &aFullName, &aJoinedAt, &aCreatedAt, &aUpdatedAt, &aTgChatID,
		&rID, &rDepID, &rSecID, &rSchedID, &rRole, &rPhone, &rFullName, &rJoinedAt, &rCreatedAt, &rUpdatedAt, &rTgChatID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	if aID.Valid {
		task.Assignee = &userDomain.User{
			ID:           aID.Int64,
			DepartmentID: aDepID.Int64,
			SectionID:    aSecID.Int64,
			Role:         int(aRole.Int64),
			Phone:        aPhone.String,
			FullName:     aFullName.String,
			JoinedAt:     aJoinedAt.Int64,
			CreatedAt:    aCreatedAt.Int64,
			UpdatedAt:    aUpdatedAt.Int64,
			TgChatID:     aTgChatID.Int64,
		}
		if aSchedID.Valid {
			task.Assignee.ScheduleID = &aSchedID.Int64
		}
	}

	if rID.Valid {
		task.Reviewer = &userDomain.User{
			ID:           rID.Int64,
			DepartmentID: rDepID.Int64,
			SectionID:    rSecID.Int64,
			Role:         int(rRole.Int64),
			Phone:        rPhone.String,
			FullName:     rFullName.String,
			JoinedAt:     rJoinedAt.Int64,
			CreatedAt:    rCreatedAt.Int64,
			UpdatedAt:    rUpdatedAt.Int64,
			TgChatID:     rTgChatID.Int64,
		}
		if rSchedID.Valid {
			task.Reviewer.ScheduleID = &rSchedID.Int64
		}
	}

	return task, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	query := `
		SELECT 
			t.id, t.parent_id, t.section_id, t.assignee_id, t.reviewer_id, t.title, t.description, t.priority, t.deadline, t.status, t.created_at, t.updated_at,
			ua.id, ua.department_id, ua.section_id, ua.schedule_id, ua.role, ua.phone, ua.full_name, ua.joined_at, ua.created_at, ua.updated_at, ua.tg_chat_id,
			ur.id, ur.department_id, ur.section_id, ur.schedule_id, ur.role, ur.phone, ur.full_name, ur.joined_at, ur.created_at, ur.updated_at, ur.tg_chat_id
		FROM tasks t
		LEFT JOIN users ua ON t.assignee_id = ua.id
		LEFT JOIN users ur ON t.reviewer_id = ur.id
		ORDER BY t.id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		var aID, aDepID, aSecID, aRole, aJoinedAt, aCreatedAt, aUpdatedAt, aTgChatID pgtype.Int8
		var aSchedID pgtype.Int8
		var aPhone, aFullName pgtype.Text
		var rID, rDepID, rSecID, rRole, rJoinedAt, rCreatedAt, rUpdatedAt, rTgChatID pgtype.Int8
		var rSchedID pgtype.Int8
		var rPhone, rFullName pgtype.Text

		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.SectionID,
			&task.AssigneeID,
			&task.ReviewerID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&aID, &aDepID, &aSecID, &aSchedID, &aRole, &aPhone, &aFullName, &aJoinedAt, &aCreatedAt, &aUpdatedAt, &aTgChatID,
			&rID, &rDepID, &rSecID, &rSchedID, &rRole, &rPhone, &rFullName, &rJoinedAt, &rCreatedAt, &rUpdatedAt, &rTgChatID,
		)
		if err != nil {
			return nil, err
		}

		if aID.Valid {
			task.Assignee = &userDomain.User{
				ID:           aID.Int64,
				DepartmentID: aDepID.Int64,
				SectionID:    aSecID.Int64,
				Role:         int(aRole.Int64),
				Phone:        aPhone.String,
				FullName:     aFullName.String,
				JoinedAt:     aJoinedAt.Int64,
				CreatedAt:    aCreatedAt.Int64,
				UpdatedAt:    aUpdatedAt.Int64,
				TgChatID:     aTgChatID.Int64,
			}
			if aSchedID.Valid {
				task.Assignee.ScheduleID = &aSchedID.Int64
			}
		}

		if rID.Valid {
			task.Reviewer = &userDomain.User{
				ID:           rID.Int64,
				DepartmentID: rDepID.Int64,
				SectionID:    rSecID.Int64,
				Role:         int(rRole.Int64),
				Phone:        rPhone.String,
				FullName:     rFullName.String,
				JoinedAt:     rJoinedAt.Int64,
				CreatedAt:    rCreatedAt.Int64,
				UpdatedAt:    rUpdatedAt.Int64,
				TgChatID:     rTgChatID.Int64,
			}
			if rSchedID.Valid {
				task.Reviewer.ScheduleID = &rSchedID.Int64
			}
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTasksWithFilter(ctx context.Context, filter *domain.TaskFilter) (*domain.TaskListResult, error) {
	baseQuery := `FROM tasks t 
	LEFT JOIN users ua ON t.assignee_id = ua.id
	LEFT JOIN users ur ON t.reviewer_id = ur.id
	WHERE 1=1`
	args := []interface{}{}
	argCount := 0

	if filter.AssigneeID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.assignee_id = $%d", argCount)
		args = append(args, *filter.AssigneeID)
	}

	if filter.ReviewerID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.reviewer_id = $%d", argCount)
		args = append(args, *filter.ReviewerID)
	}

	if filter.SectionID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.section_id = $%d", argCount)
		args = append(args, *filter.SectionID)
	}

	if filter.Status != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.status = $%d", argCount)
		args = append(args, *filter.Status)
	}

	if filter.Priority != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.priority = $%d", argCount)
		args = append(args, *filter.Priority)
	}

	if filter.ParentID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND t.parent_id = $%d", argCount)
		args = append(args, *filter.ParentID)
	}

	// Get total count
	countQuery := "SELECT COUNT(*) " + baseQuery
	var totalCount int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// Get paginated results
	selectQuery := `SELECT 
		t.id, t.parent_id, t.section_id, t.assignee_id, t.reviewer_id, t.title, t.description, t.priority, t.deadline, t.status, t.created_at, t.updated_at,
		ua.id, ua.department_id, ua.section_id, ua.schedule_id, ua.role, ua.phone, ua.full_name, ua.joined_at, ua.created_at, ua.updated_at, ua.tg_chat_id,
		ur.id, ur.department_id, ur.section_id, ur.schedule_id, ur.role, ur.phone, ur.full_name, ur.joined_at, ur.created_at, ur.updated_at, ur.tg_chat_id
	` + baseQuery + ` ORDER BY t.created_at DESC`

	if filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		if offset < 0 {
			offset = 0
		}
		argCount++
		selectQuery += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, filter.PageSize)
		argCount++
		selectQuery += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, offset)
	}

	rows, err := r.db.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		var aID, aDepID, aSecID, aRole, aJoinedAt, aCreatedAt, aUpdatedAt, aTgChatID pgtype.Int8
		var aSchedID pgtype.Int8
		var aPhone, aFullName pgtype.Text
		var rID, rDepID, rSecID, rRole, rJoinedAt, rCreatedAt, rUpdatedAt, rTgChatID pgtype.Int8
		var rSchedID pgtype.Int8
		var rPhone, rFullName pgtype.Text

		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.SectionID,
			&task.AssigneeID,
			&task.ReviewerID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&aID, &aDepID, &aSecID, &aSchedID, &aRole, &aPhone, &aFullName, &aJoinedAt, &aCreatedAt, &aUpdatedAt, &aTgChatID,
			&rID, &rDepID, &rSecID, &rSchedID, &rRole, &rPhone, &rFullName, &rJoinedAt, &rCreatedAt, &rUpdatedAt, &rTgChatID,
		)
		if err != nil {
			return nil, err
		}

		if aID.Valid {
			task.Assignee = &userDomain.User{
				ID:           aID.Int64,
				DepartmentID: aDepID.Int64,
				SectionID:    aSecID.Int64,
				Role:         int(aRole.Int64),
				Phone:        aPhone.String,
				FullName:     aFullName.String,
				JoinedAt:     aJoinedAt.Int64,
				CreatedAt:    aCreatedAt.Int64,
				UpdatedAt:    aUpdatedAt.Int64,
				TgChatID:     aTgChatID.Int64,
			}
			if aSchedID.Valid {
				task.Assignee.ScheduleID = &aSchedID.Int64
			}
		}

		if rID.Valid {
			task.Reviewer = &userDomain.User{
				ID:           rID.Int64,
				DepartmentID: rDepID.Int64,
				SectionID:    rSecID.Int64,
				Role:         int(rRole.Int64),
				Phone:        rPhone.String,
				FullName:     rFullName.String,
				JoinedAt:     rJoinedAt.Int64,
				CreatedAt:    rCreatedAt.Int64,
				UpdatedAt:    rUpdatedAt.Int64,
				TgChatID:     rTgChatID.Int64,
			}
			if rSchedID.Valid {
				task.Reviewer.ScheduleID = &rSchedID.Int64
			}
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &domain.TaskListResult{
		Tasks:      tasks,
		TotalCount: totalCount,
	}, nil
}

func (r *TaskRepository) GetTasksForCalendar(ctx context.Context, sectionID, assigneeID *int64, startDate, endDate string, status *int) ([]*domain.Task, error) {
	query := `
		SELECT 
			t.id, t.parent_id, t.section_id, t.assignee_id, t.reviewer_id, t.title, t.description, t.priority, t.deadline, t.status, t.created_at, t.updated_at,
			u.id, u.department_id, u.section_id, u.schedule_id, u.role, u.phone, u.full_name, u.joined_at, u.created_at, u.updated_at, u.tg_chat_id,
			(SELECT th.started_at FROM task_histories th 
			 JOIN task_statuses ts ON th.status = ts.id 
			 WHERE th.task_id = t.id AND ts.status_type = 10 
			 ORDER BY th.started_at DESC LIMIT 1) as end_date
		FROM tasks t
		LEFT JOIN users u ON t.assignee_id = u.id
		WHERE 1=1
	`
	args := []interface{}{}
	argCount := 0

	if startDate != "" {
		argCount++
		query += fmt.Sprintf(" AND t.created_at >= extract(epoch from $%d::timestamp)", argCount)
		args = append(args, startDate+" 00:00:00")
	}
	if endDate != "" {
		argCount++
		query += fmt.Sprintf(" AND t.created_at <= extract(epoch from $%d::timestamp)", argCount)
		args = append(args, endDate+" 23:59:59")
	}
	if sectionID != nil {
		argCount++
		query += fmt.Sprintf(" AND t.section_id = $%d", argCount)
		args = append(args, *sectionID)
	}
	if assigneeID != nil {
		argCount++
		query += fmt.Sprintf(" AND t.assignee_id = $%d", argCount)
		args = append(args, *assigneeID)
	}
	if status != nil {
		argCount++
		query += fmt.Sprintf(" AND t.status = $%d", argCount)
		args = append(args, *status)
	}

	query += " ORDER BY t.created_at ASC"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		var uID, uDepID, uSecID, uRole, uJoinedAt, uCreatedAt, uUpdatedAt, uTgChatID pgtype.Int8
		var uSchedID pgtype.Int8
		var uPhone, uFullName pgtype.Text
		var endDateVal pgtype.Int8

		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.SectionID,
			&task.AssigneeID,
			&task.ReviewerID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&uID, &uDepID, &uSecID, &uSchedID, &uRole, &uPhone, &uFullName, &uJoinedAt, &uCreatedAt, &uUpdatedAt, &uTgChatID,
			&endDateVal,
		)
		if err != nil {
			return nil, err
		}

		if endDateVal.Valid {
			task.EndDate = &endDateVal.Int64
		}

		if uID.Valid {
			task.Assignee = &userDomain.User{
				ID:           uID.Int64,
				DepartmentID: uDepID.Int64,
				SectionID:    uSecID.Int64,
				Role:         int(uRole.Int64),
				Phone:        uPhone.String,
				FullName:     uFullName.String,
				JoinedAt:     uJoinedAt.Int64,
				CreatedAt:    uCreatedAt.Int64,
				UpdatedAt:    uUpdatedAt.Int64,
				TgChatID:     uTgChatID.Int64,
			}
			if uSchedID.Valid {
				task.Assignee.ScheduleID = &uSchedID.Int64
			}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) GetEndDateForTask(ctx context.Context, taskID int64) (*int64, error) {
	query := `
		SELECT th.started_at 
		FROM task_histories th 
		JOIN task_statuses ts ON th.status = ts.id 
		WHERE th.task_id = $1 AND ts.status_type = 10 
		ORDER BY th.started_at DESC LIMIT 1
	`
	var endDate int64
	err := r.db.QueryRow(ctx, query, taskID).Scan(&endDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &endDate, nil
}
