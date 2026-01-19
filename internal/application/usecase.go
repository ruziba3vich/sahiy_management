package application

import (
	"github.com/ruziba3vich/sahiy_management/internal/application/auth"
	"github.com/ruziba3vich/sahiy_management/internal/application/branch"
	"github.com/ruziba3vich/sahiy_management/internal/application/department"
	"github.com/ruziba3vich/sahiy_management/internal/application/privilege"
	"github.com/ruziba3vich/sahiy_management/internal/application/schedule"
	"github.com/ruziba3vich/sahiy_management/internal/application/section"
	"github.com/ruziba3vich/sahiy_management/internal/application/task"
	"github.com/ruziba3vich/sahiy_management/internal/application/taskhistory"
	"github.com/ruziba3vich/sahiy_management/internal/application/taskstatus"
	"github.com/ruziba3vich/sahiy_management/internal/application/user"
	"github.com/ruziba3vich/sahiy_management/internal/application/useraction"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure"
	"github.com/ruziba3vich/sahiy_management/pkg/config"
)

type Service struct {
	department  *department.Service
	section     *section.Service
	branch      *branch.Service
	schedule    *schedule.Service
	user        *user.Service
	userAction  *useraction.Service
	taskStatus  *taskstatus.Service
	task        *task.Service
	taskHistory *taskhistory.Service
	privilege   *privilege.Service
	auth        auth.Service
}

func New(repo *infrastructure.Repository, cfg *config.Config) *Service {
	return &Service{
		department:  department.NewService(repo.GetDepartment()),
		section:     section.NewService(repo.GetSection()),
		branch:      branch.NewService(repo.GetBranch()),
		schedule:    schedule.NewService(repo.GetSchedule()),
		user:        user.NewService(repo.GetUser()),
		userAction:  useraction.NewService(repo.GetUserAction()),
		taskStatus:  taskstatus.NewService(repo.GetTaskStatus()),
		task:        task.NewService(repo.GetTask()),
		taskHistory: taskhistory.NewService(repo.GetTaskHistory()),
		privilege:   privilege.NewService(repo.GetPrivilages()),
		auth: *auth.NewService(
			repo.GetUser(),
			repo.GetPrivilages(),
			cfg.JWTSecret,
			cfg.JWTExpiryHours,
		),
	}
}

func (r *Service) GetDepartment() *department.Service {
	return r.department
}

func (r *Service) GetSection() *section.Service {
	return r.section
}

func (r *Service) GetBranch() *branch.Service {
	return r.branch
}

func (r *Service) GetSchedule() *schedule.Service {
	return r.schedule
}

func (r *Service) GetUser() *user.Service {
	return r.user
}

func (r *Service) GetUserAction() *useraction.Service {
	return r.userAction
}

func (r *Service) GetTaskStatus() *taskstatus.Service {
	return r.taskStatus
}

func (r *Service) GetTaskHistory() *taskhistory.Service {
	return r.taskHistory
}

func (r *Service) GetTask() *task.Service {
	return r.task
}

func (r *Service) GetPrivilages() *privilege.Service {
	return r.privilege
}

func (r *Service) GetAuth() *auth.Service {
	return &r.auth
}
