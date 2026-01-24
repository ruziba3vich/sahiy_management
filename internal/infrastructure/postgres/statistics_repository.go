package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/statistics"
)

type StatisticsRepository struct {
	db *pgxpool.Pool
}

func NewStatisticsRepository(db *pgxpool.Pool) domain.Repository {
	return &StatisticsRepository{db: db}
}

func (r *StatisticsRepository) GetDashboardStats(ctx context.Context, fromDate, toDate int64) (*domain.DashboardStats, error) {
	stats := &domain.DashboardStats{}

	// 1. Total Employees
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&stats.TotalEmployees)
	if err != nil {
		return nil, err
	}

	// 2. Department Stats
	depQuery := `
		SELECT d.id, d.name, COUNT(u.id) 
		FROM departments d 
		LEFT JOIN users u ON d.id = u.department_id 
		GROUP BY d.id, d.name
		ORDER BY d.id
	`
	rows, err := r.db.Query(ctx, depQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ds domain.DepartmentStat
		if err := rows.Scan(&ds.ID, &ds.Name, &ds.EmployeeCount); err != nil {
			return nil, err
		}
		stats.DepartmentStats = append(stats.DepartmentStats, ds)
	}

	// 3. Total Attendance (unique presence per day)
	attendanceQuery := `
		SELECT COUNT(DISTINCT (user_id, date_trunc('day', to_timestamp(started_at))))
		FROM user_actions
		WHERE started_at >= $1 AND started_at <= $2
	`
	err = r.db.QueryRow(ctx, attendanceQuery, fromDate, toDate).Scan(&stats.TotalAttendance)
	if err != nil {
		return nil, err
	}

	// 4. Task Stats Aggregate (Complete vs In Progress)
	taskStatQuery := `
		SELECT 
			COALESCE(SUM(CASE WHEN ts.status_type = 10 THEN 1 ELSE 0 END), 0) as completed,
			COALESCE(SUM(CASE WHEN ts.status_type != 10 THEN 1 ELSE 0 END), 0) as others
		FROM tasks t
		JOIN task_statuses ts ON t.status = ts.id
		WHERE t.updated_at >= $1 AND t.updated_at <= $2
	`
	var completed, others int64
	err = r.db.QueryRow(ctx, taskStatQuery, fromDate, toDate).Scan(&completed, &others)
	if err != nil {
		return nil, err
	}

	stats.TaskStats = []domain.TaskStatusStat{
		{ID: 1, Name: "Bajarildi", Count: completed},
		{ID: 2, Name: "Jarayonda", Count: others},
	}

	return stats, nil
}
