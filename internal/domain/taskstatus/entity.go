package taskstatus

type TaskStatus struct {
	ID        int64
	Name      string
	SectionID int64
	Color     string
}

func NewTaskStatus(name string, sectionID int64, color string) *TaskStatus {
	return &TaskStatus{
		Name:      name,
		SectionID: sectionID,
		Color:     color,
	}
}

func (t *TaskStatus) Update(name string, sectionID int64, color string) {
	t.Name = name
	t.SectionID = sectionID
	t.Color = color
}
