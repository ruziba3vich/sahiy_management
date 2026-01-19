package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/useraction"
)

var ErrUserActionNotFound = errors.New("user action not found")

type UserActionRepository struct {
	db *pgxpool.Pool
}

func NewUserActionRepository(db *pgxpool.Pool) domain.Repository {
	return &UserActionRepository{db: db}
}

func (r *UserActionRepository) CreateUserAction(ctx context.Context, action *domain.UserAction) (*domain.UserAction, error) {
	query := `
		INSERT INTO user_actions (user_id, visit_branch_id, leave_branch_id, come_status, out_status, started_at, finished_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var leaveBranchID interface{}
	if action.LeaveBranchID != nil {
		leaveBranchID = *action.LeaveBranchID
	}

	var outStatus interface{}
	if action.OutStatus != nil {
		outStatus = *action.OutStatus
	}

	var finishedAt interface{}
	if action.FinishedAt != nil {
		finishedAt = *action.FinishedAt
	}

	err := r.db.QueryRow(ctx, query,
		action.UserID,
		action.VisitBranchID,
		leaveBranchID,
		action.ComeStatus,
		outStatus,
		action.StartedAt,
		finishedAt,
	).Scan(&action.ID)
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (r *UserActionRepository) UpdateUserAction(ctx context.Context, action *domain.UserAction) (*domain.UserAction, error) {
	query := `
		UPDATE user_actions
		SET user_id = $1, visit_branch_id = $2, leave_branch_id = $3, come_status = $4, out_status = $5, started_at = $6, finished_at = $7
		WHERE id = $8
	`

	var leaveBranchID interface{}
	if action.LeaveBranchID != nil {
		leaveBranchID = *action.LeaveBranchID
	}

	var outStatus interface{}
	if action.OutStatus != nil {
		outStatus = *action.OutStatus
	}

	var finishedAt interface{}
	if action.FinishedAt != nil {
		finishedAt = *action.FinishedAt
	}

	result, err := r.db.Exec(ctx, query,
		action.UserID,
		action.VisitBranchID,
		leaveBranchID,
		action.ComeStatus,
		outStatus,
		action.StartedAt,
		finishedAt,
		action.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrUserActionNotFound
	}

	return action, nil
}

func (r *UserActionRepository) DeleteUserAction(ctx context.Context, id int64) error {
	query := `DELETE FROM user_actions WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrUserActionNotFound
	}

	return nil
}

func (r *UserActionRepository) GetUserActionByID(ctx context.Context, id int64) (*domain.UserAction, error) {
	query := `
		SELECT id, user_id, visit_branch_id, leave_branch_id, come_status, out_status, started_at, finished_at
		FROM user_actions
		WHERE id = $1
	`

	action := &domain.UserAction{}
	var leaveBranchID sql.NullInt64
	var outStatus sql.NullInt64
	var finishedAt sql.NullInt64

	err := r.db.QueryRow(ctx, query, id).Scan(
		&action.ID,
		&action.UserID,
		&action.VisitBranchID,
		&leaveBranchID,
		&action.ComeStatus,
		&outStatus,
		&action.StartedAt,
		&finishedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserActionNotFound
		}
		return nil, err
	}

	if leaveBranchID.Valid {
		action.LeaveBranchID = &leaveBranchID.Int64
	}
	if outStatus.Valid {
		value := int(outStatus.Int64)
		action.OutStatus = &value
	}
	if finishedAt.Valid {
		action.FinishedAt = &finishedAt.Int64
	}

	return action, nil
}

func (r *UserActionRepository) GetAllUserActions(ctx context.Context) ([]*domain.UserAction, error) {
	query := `
		SELECT id, user_id, visit_branch_id, leave_branch_id, come_status, out_status, started_at, finished_at
		FROM user_actions
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []*domain.UserAction
	for rows.Next() {
		action := &domain.UserAction{}
		var leaveBranchID sql.NullInt64
		var outStatus sql.NullInt64
		var finishedAt sql.NullInt64

		err := rows.Scan(
			&action.ID,
			&action.UserID,
			&action.VisitBranchID,
			&leaveBranchID,
			&action.ComeStatus,
			&outStatus,
			&action.StartedAt,
			&finishedAt,
		)
		if err != nil {
			return nil, err
		}

		if leaveBranchID.Valid {
			action.LeaveBranchID = &leaveBranchID.Int64
		}
		if outStatus.Valid {
			value := int(outStatus.Int64)
			action.OutStatus = &value
		}
		if finishedAt.Valid {
			action.FinishedAt = &finishedAt.Int64
		}

		actions = append(actions, action)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}
