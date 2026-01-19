package section

import "time"

type Section struct {
	ID           int64
	DepartmentID int64
	Name         string
	CreatedAt    int64
	UpdatedAt    int64
}

func NewSection(departmentID int64, name string) *Section {
	now := time.Now().Unix()
	return &Section{
		DepartmentID: departmentID,
		Name:         name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (s *Section) Update(departmentID int64, name string) {
	s.DepartmentID = departmentID
	s.Name = name
	s.UpdatedAt = time.Now().Unix()
}
