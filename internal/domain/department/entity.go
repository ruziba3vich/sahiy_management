package department

import "time"

type Department struct {
	ID        int64
	Name      string
	Status    int
	CreatedAt int64
	UpdatedAt int64
}

func NewDepartment(name string, status int) *Department {
	now := time.Now().Unix()
	return &Department{
		Name:      name,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (d *Department) Update(name string, status int) {
	d.Name = name
	d.Status = status
	d.UpdatedAt = time.Now().Unix()
}

type DepartmentBranch struct {
	ID           int64
	BranchID     int64
	DepartmentID int64
	Status       int
	CreatedAt    int64
}
