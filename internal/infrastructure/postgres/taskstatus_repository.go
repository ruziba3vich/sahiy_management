package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
)

var ErrTaskStatusNotFound = errors.New("task status not found")

type TaskStatusRepository struct {
	db *pgxpool.Pool
}

func NewTaskStatusRepository(db *pgxpool.Pool) domain.Repository {
	return &TaskStatusRepository{db: db}
}

func (r *TaskStatusRepository) CreateTaskStatus(ctx context.Context, ts *domain.TaskStatus) (*domain.TaskStatus, error) {
	query := `
		INSERT INTO task_statuses (name, section_id, color)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query, ts.Name, ts.SectionID, ts.Color).Scan(&ts.ID)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *TaskStatusRepository) UpdateTaskStatus(ctx context.Context, ts *domain.TaskStatus) (*domain.TaskStatus, error) {
	query := `
		UPDATE task_statuses
		SET name = $1, section_id = $2, color = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(ctx, query, ts.Name, ts.SectionID, ts.Color, ts.ID)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrTaskStatusNotFound
	}

	return ts, nil
}

func (r *TaskStatusRepository) DeleteTaskStatus(ctx context.Context, id int64) error {
	query := `DELETE FROM task_statuses WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrTaskStatusNotFound
	}

	return nil
}

func (r *TaskStatusRepository) GetTaskStatusByID(ctx context.Context, id int64) (*domain.TaskStatus, error) {
	query := `
		SELECT id, name, section_id, color
		FROM task_statuses
		WHERE id = $1
	`

	ts := &domain.TaskStatus{}
	var color pgtype.Text
	err := r.db.QueryRow(ctx, query, id).Scan(&ts.ID, &ts.Name, &ts.SectionID, &color)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskStatusNotFound
		}
		return nil, err
	}
	ts.Color = color.String

	return ts, nil
}

func (r *TaskStatusRepository) GetAllTaskStatuses(ctx context.Context) ([]*domain.TaskStatus, error) {
	query := `
		SELECT id, name, section_id, color
		FROM task_statuses
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*domain.TaskStatus
	for rows.Next() {
		ts := &domain.TaskStatus{}
		var color pgtype.Text
		err := rows.Scan(&ts.ID, &ts.Name, &ts.SectionID, &color)
		if err != nil {
			return nil, err
		}
		ts.Color = color.String
		statuses = append(statuses, ts)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return statuses, nil
}
