package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/section"
)

var ErrSectionNotFound = errors.New("section not found")

type SectionRepository struct {
	db *pgxpool.Pool
}

func NewSectionRepository(db *pgxpool.Pool) domain.Repository {
	return &SectionRepository{db: db}
}

func (r *SectionRepository) CreateSection(ctx context.Context, sec *domain.Section) (*domain.Section, error) {
	query := `
		INSERT INTO sections (department_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		sec.DepartmentID,
		sec.Name,
		sec.CreatedAt,
		sec.UpdatedAt,
	).Scan(&sec.ID)
	if err != nil {
		return nil, err
	}

	return sec, nil
}

func (r *SectionRepository) UpdateSection(ctx context.Context, sec *domain.Section) (*domain.Section, error) {
	query := `
		UPDATE sections
		SET department_id = $1, name = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(ctx, query,
		sec.DepartmentID,
		sec.Name,
		sec.UpdatedAt,
		sec.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrSectionNotFound
	}

	return sec, nil
}

func (r *SectionRepository) DeleteSection(ctx context.Context, id int64) error {
	query := `DELETE FROM sections WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrSectionNotFound
	}

	return nil
}

func (r *SectionRepository) GetSectionByID(ctx context.Context, id int64) (*domain.Section, error) {
	query := `
		SELECT id, department_id, name, created_at, updated_at
		FROM sections
		WHERE id = $1
	`

	sec := &domain.Section{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&sec.ID,
		&sec.DepartmentID,
		&sec.Name,
		&sec.CreatedAt,
		&sec.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSectionNotFound
		}
		return nil, err
	}

	return sec, nil
}

func (r *SectionRepository) GetAllSections(ctx context.Context) ([]*domain.Section, error) {
	query := `
		SELECT id, department_id, name, created_at, updated_at
		FROM sections
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []*domain.Section
	for rows.Next() {
		sec := &domain.Section{}
		err := rows.Scan(
			&sec.ID,
			&sec.DepartmentID,
			&sec.Name,
			&sec.CreatedAt,
			&sec.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sections = append(sections, sec)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}
