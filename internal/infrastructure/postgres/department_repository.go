package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/department"
)

var (
	ErrDepartmentNotFound        = errors.New("Department not found")
	ErrDepartmentBranchNotFound  = errors.New("Department branch not found")
	ErrDuplicateDepartmentBranch = errors.New("Branch already attached to this department")
)

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

func (r *DepartmentRepository) AttachBranch(ctx context.Context, db *domain.DepartmentBranch) (*domain.DepartmentBranch, error) {
	query := `
		INSERT INTO department_branches (branch_id, department_id, status, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		db.BranchID,
		db.DepartmentID,
		db.Status,
		db.CreatedAt,
	).Scan(&db.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrDuplicateDepartmentBranch
		}
		return nil, err
	}

	return db, nil
}

func (r *DepartmentRepository) UpdateDepartmentBranchStatus(ctx context.Context, id int64, status int) error {
	query := `UPDATE department_branches SET status = $1 WHERE id = $2`

	result, err := r.db.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrDepartmentBranchNotFound
	}

	return nil
}

func (r *DepartmentRepository) GetBranchesByDepartmentID(ctx context.Context, departmentID int64) ([]*domain.DepartmentBranch, error) {
	query := `
		SELECT id, branch_id, department_id, status, created_at
		FROM department_branches
		WHERE department_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []*domain.DepartmentBranch
	for rows.Next() {
		db := &domain.DepartmentBranch{}
		err := rows.Scan(
			&db.ID,
			&db.BranchID,
			&db.DepartmentID,
			&db.Status,
			&db.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		branches = append(branches, db)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return branches, nil
}
