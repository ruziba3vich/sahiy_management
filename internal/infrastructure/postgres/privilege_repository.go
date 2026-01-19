package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/privilege"
)

var ErrPrivilegeNotFound = errors.New("privilege not found")
var ErrUserPrivilegeNotFound = errors.New("user privilege not found")

type PrivilegeRepository struct {
	db *pgxpool.Pool
}

func NewPrivilegeRepository(db *pgxpool.Pool) domain.Repository {
	return &PrivilegeRepository{db: db}
}

func (r *PrivilegeRepository) CreatePrivilege(ctx context.Context, p *domain.Privilege) (*domain.Privilege, error) {
	query := `
		INSERT INTO privileges (name, resource, action, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		p.Name,
		p.Resource,
		p.Action,
		p.Description,
		p.CreatedAt,
		p.UpdatedAt,
	).Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PrivilegeRepository) UpdatePrivilege(ctx context.Context, p *domain.Privilege) (*domain.Privilege, error) {
	query := `
		UPDATE privileges
		SET name = $1, resource = $2, action = $3, description = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(ctx, query,
		p.Name,
		p.Resource,
		p.Action,
		p.Description,
		p.UpdatedAt,
		p.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrPrivilegeNotFound
	}

	return p, nil
}

func (r *PrivilegeRepository) DeletePrivilege(ctx context.Context, id int64) error {
	query := `DELETE FROM privileges WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrPrivilegeNotFound
	}

	return nil
}

func (r *PrivilegeRepository) GetPrivilegeByID(ctx context.Context, id int64) (*domain.Privilege, error) {
	query := `
		SELECT id, name, resource, action, description, created_at, updated_at
		FROM privileges
		WHERE id = $1
	`

	p := &domain.Privilege{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Resource,
		&p.Action,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPrivilegeNotFound
		}
		return nil, err
	}

	return p, nil
}

func (r *PrivilegeRepository) GetAllPrivileges(ctx context.Context) ([]*domain.Privilege, error) {
	query := `
		SELECT id, name, resource, action, description, created_at, updated_at
		FROM privileges
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privileges []*domain.Privilege
	for rows.Next() {
		p := &domain.Privilege{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Resource,
			&p.Action,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		privileges = append(privileges, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return privileges, nil
}

func (r *PrivilegeRepository) GetPrivilegeByResourceAction(ctx context.Context, resource, action string) (*domain.Privilege, error) {
	query := `
		SELECT id, name, resource, action, description, created_at, updated_at
		FROM privileges
		WHERE resource = $1 AND action = $2
	`

	p := &domain.Privilege{}
	err := r.db.QueryRow(ctx, query, resource, action).Scan(
		&p.ID,
		&p.Name,
		&p.Resource,
		&p.Action,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPrivilegeNotFound
		}
		return nil, err
	}

	return p, nil
}

func (r *PrivilegeRepository) AssignPrivilegeToUser(ctx context.Context, up *domain.UserPrivilege) (*domain.UserPrivilege, error) {
	query := `
		INSERT INTO user_privileges (user_id, privilege_id, granted_by, granted_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		up.UserID,
		up.PrivilegeID,
		up.GrantedBy,
		up.GrantedAt,
	).Scan(&up.ID)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (r *PrivilegeRepository) RevokePrivilegeFromUser(ctx context.Context, userID, privilegeID int64) error {
	query := `DELETE FROM user_privileges WHERE user_id = $1 AND privilege_id = $2`

	result, err := r.db.Exec(ctx, query, userID, privilegeID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrUserPrivilegeNotFound
	}

	return nil
}

func (r *PrivilegeRepository) GetUserPrivileges(ctx context.Context, userID int64) ([]*domain.Privilege, error) {
	query := `
		SELECT p.id, p.name, p.resource, p.action, p.description, p.created_at, p.updated_at
		FROM privileges p
		INNER JOIN user_privileges up ON p.id = up.privilege_id
		WHERE up.user_id = $1
		ORDER BY p.id
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privileges []*domain.Privilege
	for rows.Next() {
		p := &domain.Privilege{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Resource,
			&p.Action,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		privileges = append(privileges, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return privileges, nil
}

func (r *PrivilegeRepository) HasPrivilege(ctx context.Context, userID int64, resource, action string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM user_privileges up
			INNER JOIN privileges p ON up.privilege_id = p.id
			WHERE up.user_id = $1 AND p.resource = $2 AND p.action = $3
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, userID, resource, action).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
