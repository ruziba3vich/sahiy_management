package task

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Task struct {
	ID          int64
	ParentID    pgtype.Int8
	SectionID   int64
	UserID      pgtype.Int8
	Title       string
	Description string
	Priority    int
	Deadline    int64
	Status      int
	CreatedAt   int64
	UpdatedAt   int64
}

func NewTask(parentID pgtype.Int8, sectionID int64, userID pgtype.Int8, title, description string, priority int, deadline int64, status int) *Task {
	now := time.Now().Unix()
	return &Task{
		ParentID:    parentID,
		SectionID:   sectionID,
		UserID:      userID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Deadline:    deadline,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) Update(parentID pgtype.Int8, sectionID int64, userID pgtype.Int8, title, description string, priority int, deadline int64, status int) {
	t.ParentID = parentID
	t.SectionID = sectionID
	t.UserID = userID
	t.Title = title
	t.Description = description
	t.Priority = priority
	t.Deadline = deadline
	t.Status = status
	t.UpdatedAt = time.Now().Unix()
}
