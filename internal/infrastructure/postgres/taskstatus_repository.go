package postgres

import (
	"context"
	"database/sql"
	"errors"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
)

var ErrTaskStatusNotFound = errors.New("task status not found")

type TaskStatusRepository struct {
	db *sql.DB
}

func NewTaskStatusRepository(db *sql.DB) domain.Repository {
	return &TaskStatusRepository{db: db}
}

func (r *TaskStatusRepository) CreateTaskStatus(ctx context.Context, ts *domain.TaskStatus) (*domain.TaskStatus, error) {
	query := `
		INSERT INTO task_statuses (name, type)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, ts.Name, ts.Type).Scan(&ts.ID)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *TaskStatusRepository) UpdateTaskStatus(ctx context.Context, ts *domain.TaskStatus) (*domain.TaskStatus, error) {
	query := `
		UPDATE task_statuses
		SET name = $1, type = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, ts.Name, ts.Type, ts.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrTaskStatusNotFound
	}

	return ts, nil
}

func (r *TaskStatusRepository) DeleteTaskStatus(ctx context.Context, id int64) error {
	query := `DELETE FROM task_statuses WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrTaskStatusNotFound
	}

	return nil
}

func (r *TaskStatusRepository) GetTaskStatusByID(ctx context.Context, id int64) (*domain.TaskStatus, error) {
	query := `
		SELECT id, name, type
		FROM task_statuses
		WHERE id = $1
	`

	ts := &domain.TaskStatus{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ts.ID, &ts.Name, &ts.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTaskStatusNotFound
		}
		return nil, err
	}

	return ts, nil
}

func (r *TaskStatusRepository) GetAllTaskStatuses(ctx context.Context) ([]*domain.TaskStatus, error) {
	query := `
		SELECT id, name, type
		FROM task_statuses
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*domain.TaskStatus
	for rows.Next() {
		ts := &domain.TaskStatus{}
		err := rows.Scan(&ts.ID, &ts.Name, &ts.Type)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, ts)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return statuses, nil
}
