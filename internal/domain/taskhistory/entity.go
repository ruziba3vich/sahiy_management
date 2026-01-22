package taskhistory

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type TaskHistory struct {
	ID           int64
	TaskID       int64
	UserID       int64
	Status       int
	StartedAt    int64
	FinishedAt   pgtype.Int8
	UserFullName string
	StatusName   string
	Color        string
}

func NewTaskHistory(taskID, userID int64, status int) *TaskHistory {
	return &TaskHistory{
		TaskID:    taskID,
		UserID:    userID,
		Status:    status,
		StartedAt: time.Now().Unix(),
	}
}

func (th *TaskHistory) Update(taskID, userID int64, status int, finishedAt *int64) {
	th.TaskID = taskID
	th.UserID = userID
	th.Status = status
	if finishedAt != nil {
		th.FinishedAt = pgtype.Int8{Int64: *finishedAt, Valid: true}
	} else {
		th.FinishedAt = pgtype.Int8{Valid: false}
	}
}

func (th *TaskHistory) Finish() {
	th.FinishedAt = pgtype.Int8{Int64: time.Now().Unix(), Valid: true}
}
