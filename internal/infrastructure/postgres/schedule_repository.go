package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/schedule"
)

var ErrScheduleNotFound = errors.New("schedule not found")

type ScheduleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRepository(db *pgxpool.Pool) domain.Repository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) CreateSchedule(ctx context.Context, schedule *domain.Schedule) (*domain.Schedule, error) {
	query := `
		INSERT INTO schedules (name, timezone, week_days, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		schedule.Name,
		schedule.Timezone,
		schedule.WeekDays,
		schedule.Status,
		schedule.CreatedAt,
		schedule.UpdatedAt,
	).Scan(&schedule.ID)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *ScheduleRepository) UpdateSchedule(ctx context.Context, schedule *domain.Schedule) (*domain.Schedule, error) {
	query := `
		UPDATE schedules
		SET name = $1, timezone = $2, week_days = $3, status = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(ctx, query,
		schedule.Name,
		schedule.Timezone,
		schedule.WeekDays,
		schedule.Status,
		schedule.UpdatedAt,
		schedule.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrScheduleNotFound
	}

	return schedule, nil
}

func (r *ScheduleRepository) DeleteSchedule(ctx context.Context, id int64) error {
	query := `DELETE FROM schedules WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrScheduleNotFound
	}

	return nil
}

func (r *ScheduleRepository) GetScheduleByID(ctx context.Context, id int64) (*domain.Schedule, error) {
	query := `
		SELECT id, name, timezone, week_days, status, created_at, updated_at
		FROM schedules
		WHERE id = $1
	`

	schedule := &domain.Schedule{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&schedule.ID,
		&schedule.Name,
		&schedule.Timezone,
		&schedule.WeekDays,
		&schedule.Status,
		&schedule.CreatedAt,
		&schedule.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrScheduleNotFound
		}
		return nil, err
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetAllSchedules(ctx context.Context) ([]*domain.Schedule, error) {
	query := `
		SELECT id, name, timezone, week_days, status, created_at, updated_at
		FROM schedules
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*domain.Schedule
	for rows.Next() {
		schedule := &domain.Schedule{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.Name,
			&schedule.Timezone,
			&schedule.WeekDays,
			&schedule.Status,
			&schedule.CreatedAt,
			&schedule.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}
