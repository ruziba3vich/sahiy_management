package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
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
		INSERT INTO tasks (parent_id, section_id, user_id, title, description, priority, deadline, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		task.ParentID,
		task.SectionID,
		task.UserID,
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
		SET parent_id = $1, section_id = $2, user_id = $3, title = $4, description = $5, priority = $6, deadline = $7, status = $8, updated_at = $9
		WHERE id = $10
	`

	result, err := r.db.Exec(ctx, query,
		task.ParentID,
		task.SectionID,
		task.UserID,
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
		SELECT id, parent_id, section_id, user_id, title, description, priority, deadline, status, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	task := &domain.Task{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.ParentID,
		&task.SectionID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	query := `
		SELECT id, parent_id, section_id, user_id, title, description, priority, deadline, status, created_at, updated_at
		FROM tasks
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.SectionID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTasksWithFilter(ctx context.Context, filter *domain.TaskFilter) (*domain.TaskListResult, error) {
	baseQuery := `FROM tasks WHERE 1=1`
	args := []interface{}{}
	argCount := 0

	if filter.UserID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND user_id = $%d", argCount)
		args = append(args, *filter.UserID)
	}

	if filter.SectionID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND section_id = $%d", argCount)
		args = append(args, *filter.SectionID)
	}

	if filter.Status != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *filter.Status)
	}

	if filter.Priority != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND priority = $%d", argCount)
		args = append(args, *filter.Priority)
	}

	if filter.ParentID != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND parent_id = $%d", argCount)
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
	selectQuery := `SELECT id, parent_id, section_id, user_id, title, description, priority, deadline, status, created_at, updated_at ` + baseQuery + ` ORDER BY created_at DESC`

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
		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.SectionID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
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
