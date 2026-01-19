package postgres

import (
	"context"
	"errors"

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
		INSERT INTO tasks (parent_id, section_id, title, description, priority, deadline, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		task.ParentID,
		task.SectionID,
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
		SET parent_id = $1, section_id = $2, title = $3, description = $4, priority = $5, deadline = $6, status = $7, updated_at = $8
		WHERE id = $9
	`

	result, err := r.db.Exec(ctx, query,
		task.ParentID,
		task.SectionID,
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
		SELECT id, parent_id, section_id, title, description, priority, deadline, status, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	task := &domain.Task{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.ParentID,
		&task.SectionID,
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
		SELECT id, parent_id, section_id, title, description, priority, deadline, status, created_at, updated_at
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
