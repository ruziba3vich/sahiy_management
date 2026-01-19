package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/department"
)

var ErrDepartmentNotFound = errors.New("department not found")

type DepartmentRepository struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(db *pgxpool.Pool) domain.Repository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) CreateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error) {
	query := `
		INSERT INTO departments (name, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		dept.Name,
		dept.Status,
		dept.CreatedAt,
		dept.UpdatedAt,
	).Scan(&dept.ID)
	if err != nil {
		return nil, err
	}

	return dept, nil
}

func (r *DepartmentRepository) UpdateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error) {
	query := `
		UPDATE departments
		SET name = $1, status = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(ctx, query,
		dept.Name,
		dept.Status,
		dept.UpdatedAt,
		dept.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrDepartmentNotFound
	}

	return dept, nil
}

func (r *DepartmentRepository) DeleteDepartment(ctx context.Context, id int64) error {
	query := `DELETE FROM departments WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrDepartmentNotFound
	}

	return nil
}

func (r *DepartmentRepository) GetDepartmentByID(ctx context.Context, id int64) (*domain.Department, error) {
	query := `
		SELECT id, name, status, created_at, updated_at
		FROM departments
		WHERE id = $1
	`

	dept := &domain.Department{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&dept.ID,
		&dept.Name,
		&dept.Status,
		&dept.CreatedAt,
		&dept.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDepartmentNotFound
		}
		return nil, err
	}

	return dept, nil
}

func (r *DepartmentRepository) GetAllDepartments(ctx context.Context) ([]*domain.Department, error) {
	query := `
		SELECT id, name, status, created_at, updated_at
		FROM departments
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*domain.Department
	for rows.Next() {
		dept := &domain.Department{}
		err := rows.Scan(
			&dept.ID,
			&dept.Name,
			&dept.Status,
			&dept.CreatedAt,
			&dept.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}
