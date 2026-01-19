package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ruziba3vich/sahiy_management/internal/domain/department"
	"github.com/ruziba3vich/sahiy_management/internal/domain/privilege"
	"github.com/ruziba3vich/sahiy_management/internal/domain/section"
	"github.com/ruziba3vich/sahiy_management/internal/domain/task"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
	"github.com/ruziba3vich/sahiy_management/internal/domain/user"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
)

type Repository struct {
	department  department.Repository
	section     section.Repository
	user        user.Repository
	taskStatus  taskstatus.Repository
	taskHistory taskhistory.Repository
	tasks       task.Repository
	privilages  privilege.Repository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		department:  postgres.NewDepartmentRepository(db),
		section:     postgres.NewSectionRepository(db),
		user:        postgres.NewUserRepository(db),
		taskStatus:  postgres.NewTaskStatusRepository(db),
		tasks:       postgres.NewTaskRepository(db),
		taskHistory: postgres.NewTaskHistoryRepository(db),
		privilages:  postgres.NewPrivilegeRepository(db),
	}
}

func (r *Repository) GetDepartment() department.Repository {
	return r.department
}

func (r *Repository) GetSection() section.Repository {
	return r.section
}

func (r *Repository) GetUser() user.Repository {
	return r.user
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

func (r *Repository) GetPrivilages() privilege.Repository {
	return r.privilages
}
