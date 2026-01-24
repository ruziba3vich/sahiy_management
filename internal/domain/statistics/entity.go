package statistics

type DashboardStats struct {
	TotalEmployees  int64            `json:"total_employees"`
	DepartmentStats []DepartmentStat `json:"department_stats"`
	TotalAttendance int64            `json:"total_attendance"`
	TaskStats       []TaskStatusStat `json:"task_stats"`
}

type DepartmentStat struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	EmployeeCount int64  `json:"employee_count"`
}

type TaskStatusStat struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
