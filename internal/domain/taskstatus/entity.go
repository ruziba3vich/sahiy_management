package taskstatus

type TaskStatus struct {
	ID   int64
	Name string
	Type int
}

func NewTaskStatus(name string, statusType int) *TaskStatus {
	return &TaskStatus{
		Name: name,
		Type: statusType,
	}
}

func (t *TaskStatus) Update(name string, statusType int) {
	t.Name = name
	t.Type = statusType
}
