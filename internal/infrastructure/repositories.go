package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ruziba3vich/sahiy_management/internal/domain/branch"
	"github.com/ruziba3vich/sahiy_management/internal/domain/department"
	"github.com/ruziba3vich/sahiy_management/internal/domain/schedule"
	"github.com/ruziba3vich/sahiy_management/internal/domain/section"
	"github.com/ruziba3vich/sahiy_management/internal/domain/statistics"
	"github.com/ruziba3vich/sahiy_management/internal/domain/task"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
	"github.com/ruziba3vich/sahiy_management/internal/domain/user"
	"github.com/ruziba3vich/sahiy_management/internal/domain/useraction"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
)

type Repository struct {
	department  department.Repository
	section     section.Repository
	branch      branch.Repository
	schedule    schedule.Repository
	user        user.Repository
	userAction  useraction.Repository
	taskStatus  taskstatus.Repository
	taskHistory taskhistory.Repository
	tasks       task.Repository
	statistics  statistics.Repository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		department:  postgres.NewDepartmentRepository(db),
		section:     postgres.NewSectionRepository(db),
		branch:      postgres.NewBranchRepository(db),
		schedule:    postgres.NewScheduleRepository(db),
		user:        postgres.NewUserRepository(db),
		userAction:  postgres.NewUserActionRepository(db),
		taskStatus:  postgres.NewTaskStatusRepository(db),
		tasks:       postgres.NewTaskRepository(db),
		taskHistory: postgres.NewTaskHistoryRepository(db),
		statistics:  postgres.NewStatisticsRepository(db),
	}
}

func (r *Repository) GetDepartment() department.Repository {
	return r.department
}

func (r *Repository) GetSection() section.Repository {
	return r.section
}

func (r *Repository) GetBranch() branch.Repository {
	return r.branch
}

func (r *Repository) GetSchedule() schedule.Repository {
	return r.schedule
}

func (r *Repository) GetUser() user.Repository {
	return r.user
}

func (r *Repository) GetUserAction() useraction.Repository {
	return r.userAction
}

func (r *Repository) GetTaskStatus() taskstatus.Repository {
	return r.taskStatus
}

func (r *Repository) GetTaskHistory() taskhistory.Repository {
	return r.taskHistory
}

func (r *Repository) GetTask() task.Repository {
	return r.tasks
}

func (r *Repository) GetStatistics() statistics.Repository {
	return r.statistics
}
