package taskstatus

type TaskStatus struct {
	ID         int64
	Name       string
	SectionID  int64
	Color      string
	StatusType int64
	Sort       int
}

func NewTaskStatus(name string, sectionID int64, color string, status_type int64, sort int) *TaskStatus {
	return &TaskStatus{
		Name:       name,
		SectionID:  sectionID,
		Color:      color,
		StatusType: status_type,
		Sort:       sort,
	}
}

func (t *TaskStatus) Update(name string, sectionID int64, color string, status_type int64, sort int) {
	t.Name = name
	t.SectionID = sectionID
	t.Color = color
	t.StatusType = status_type
	t.Sort = sort
}
