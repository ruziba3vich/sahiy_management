package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/user"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (department_id, section_id, schedule_id, role, phone, full_name, joined_at, created_at, updated_at, tg_chat_id, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		user.DepartmentID,
		user.SectionID,
		user.ScheduleID,
		user.Role,
		user.Phone,
		user.FullName,
		user.JoinedAt,
		user.CreatedAt,
		user.UpdatedAt,
		user.TgChatID,
		user.PasswordHash,
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET department_id = $1, section_id = $2, schedule_id = $3, role = $4, phone = $5, full_name = $6, updated_at = $7, tg_chat_id = $8
		WHERE id = $9
	`

	result, err := r.db.Exec(ctx, query,
		user.DepartmentID,
		user.SectionID,
		user.ScheduleID,
		user.Role,
		user.Phone,
		user.FullName,
		user.UpdatedAt,
		user.TgChatID,
		user.ID,
	)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT id, department_id, section_id, schedule_id, role, phone, full_name, joined_at, created_at, updated_at, tg_chat_id, password_hash
		FROM users
		WHERE id = $1
	`

	user := &domain.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.DepartmentID,
		&user.SectionID,
		&user.ScheduleID,
		&user.Role,
		&user.Phone,
		&user.FullName,
		&user.JoinedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.TgChatID,
		&user.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := `
		SELECT id, department_id, section_id, schedule_id, role, phone, full_name, joined_at, created_at, updated_at, tg_chat_id, password_hash
		FROM users
		WHERE phone = $1
	`

	user := &domain.User{}
	err := r.db.QueryRow(ctx, query, phone).Scan(
		&user.ID,
		&user.DepartmentID,
		&user.SectionID,
		&user.ScheduleID,
		&user.Role,
		&user.Phone,
		&user.FullName,
		&user.JoinedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.TgChatID,
		&user.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	query := `
		SELECT id, department_id, section_id, schedule_id, role, phone, full_name, joined_at, created_at, updated_at, tg_chat_id, password_hash
		FROM users
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.DepartmentID,
			&user.SectionID,
			&user.ScheduleID,
			&user.Role,
			&user.Phone,
			&user.FullName,
			&user.JoinedAt,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.TgChatID,
			&user.PasswordHash,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
