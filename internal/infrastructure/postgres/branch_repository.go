package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/branch"
)

var ErrBranchNotFound = errors.New("branch not found")

type BranchRepository struct {
	db *pgxpool.Pool
}

func NewBranchRepository(db *pgxpool.Pool) domain.Repository {
	return &BranchRepository{db: db}
}

func (r *BranchRepository) CreateBranch(ctx context.Context, branch *domain.Branch) (*domain.Branch, error) {
	query := `
		INSERT INTO branchs (name, lat, long, radius, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		branch.Name,
		branch.Lat,
		branch.Long,
		branch.Radius,
		branch.Status,
		branch.CreatedAt,
		branch.UpdatedAt,
	).Scan(&branch.ID)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (r *BranchRepository) UpdateBranch(ctx context.Context, branch *domain.Branch) (*domain.Branch, error) {
	query := `
		UPDATE branchs
		SET name = $1, lat = $2, long = $3, radius = $4, status = $5, updated_at = $6
		WHERE id = $7
	`

	result, err := r.db.Exec(ctx, query,
		branch.Name,
		branch.Lat,
		branch.Long,
		branch.Radius,
		branch.Status,
		branch.UpdatedAt,
		branch.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrBranchNotFound
	}

	return branch, nil
}

func (r *BranchRepository) DeleteBranch(ctx context.Context, id int64) error {
	query := `DELETE FROM branchs WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrBranchNotFound
	}

	return nil
}

func (r *BranchRepository) GetBranchByID(ctx context.Context, id int64) (*domain.Branch, error) {
	query := `
		SELECT id, name, lat, long, radius, status, created_at, updated_at
		FROM branchs
		WHERE id = $1
	`

	branch := &domain.Branch{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Lat,
		&branch.Long,
		&branch.Radius,
		&branch.Status,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrBranchNotFound
		}
		return nil, err
	}

	return branch, nil
}

func (r *BranchRepository) GetAllBranches(ctx context.Context) ([]*domain.Branch, error) {
	query := `
		SELECT id, name, lat, long, radius, status, created_at, updated_at
		FROM branchs
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []*domain.Branch
	for rows.Next() {
		branch := &domain.Branch{}
		err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Lat,
			&branch.Long,
			&branch.Radius,
			&branch.Status,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return branches, nil
}
