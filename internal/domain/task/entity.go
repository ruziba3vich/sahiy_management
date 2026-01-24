package task

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ruziba3vich/sahiy_management/internal/domain/user"
)

type Task struct {
	ID          int64
	ParentID    pgtype.Int8
	SectionID   int64
	AssigneeID  pgtype.Int8
	Assignee    *user.User
	ReviewerID  pgtype.Int8
	Reviewer    *user.User
	Title       string
	Description string
	Priority    int
	Deadline    pgtype.Int8
	Status      int
	CreatedAt   int64
	UpdatedAt   int64
	EndDate     *int64
}

func NewTask(parentID pgtype.Int8, sectionID int64, assigneeID, reviewerID pgtype.Int8, title, description string, priority int, deadline pgtype.Int8, status int) *Task {
	now := time.Now().Unix()
	return &Task{
		ParentID:    parentID,
		SectionID:   sectionID,
		AssigneeID:  assigneeID,
		ReviewerID:  reviewerID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Deadline:    deadline,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) Update(parentID pgtype.Int8, sectionID int64, assigneeID, reviewerID pgtype.Int8, title, description string, priority int, deadline pgtype.Int8, status int) {
	t.ParentID = parentID
	t.SectionID = sectionID
	t.AssigneeID = assigneeID
	t.ReviewerID = reviewerID
	t.Title = title
	t.Description = description
	t.Priority = priority
	t.Deadline = deadline
	t.Status = status
	t.UpdatedAt = time.Now().Unix()
}
