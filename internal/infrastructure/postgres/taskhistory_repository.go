package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
)

var ErrTaskHistoryNotFound = errors.New("task history not found")

type TaskHistoryRepository struct {
	db *pgxpool.Pool
}

func NewTaskHistoryRepository(db *pgxpool.Pool) domain.Repository {
	return &TaskHistoryRepository{db: db}
}

func (r *TaskHistoryRepository) CreateTaskHistory(ctx context.Context, th *domain.TaskHistory) (*domain.TaskHistory, error) {
	query := `
		INSERT INTO task_histories (task_id, user_id, status, started_at, finished_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		th.TaskID,
		th.UserID,
		th.Status,
		th.StartedAt,
		th.FinishedAt,
	).Scan(&th.ID)
	if err != nil {
		return nil, err
	}

	return th, nil
}

func (r *TaskHistoryRepository) UpdateTaskHistory(ctx context.Context, th *domain.TaskHistory) (*domain.TaskHistory, error) {
	query := `
		UPDATE task_histories
		SET task_id = $1, user_id = $2, status = $3, started_at = $4, finished_at = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(ctx, query,
		th.TaskID,
		th.UserID,
		th.Status,
		th.StartedAt,
		th.FinishedAt,
		th.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrTaskHistoryNotFound
	}

	return th, nil
}

func (r *TaskHistoryRepository) DeleteTaskHistory(ctx context.Context, id int64) error {
	query := `DELETE FROM task_histories WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrTaskHistoryNotFound
	}

	return nil
}

func (r *TaskHistoryRepository) GetTaskHistoryByID(ctx context.Context, id int64) (*domain.TaskHistory, error) {
	query := `
		SELECT id, task_id, user_id, status, started_at, finished_at
		FROM task_histories
		WHERE id = $1
	`

	th := &domain.TaskHistory{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&th.ID,
		&th.TaskID,
		&th.UserID,
		&th.Status,
		&th.StartedAt,
		&th.FinishedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskHistoryNotFound
		}
		return nil, err
	}

	return th, nil
}

func (r *TaskHistoryRepository) GetAllTaskHistories(ctx context.Context) ([]*domain.TaskHistory, error) {
	query := `
		SELECT id, task_id, user_id, status, started_at, finished_at
		FROM task_histories
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*domain.TaskHistory
	for rows.Next() {
		th := &domain.TaskHistory{}
		err := rows.Scan(
			&th.ID,
			&th.TaskID,
			&th.UserID,
			&th.Status,
			&th.StartedAt,
			&th.FinishedAt,
		)
		if err != nil {
			return nil, err
		}
		histories = append(histories, th)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return histories, nil
}
