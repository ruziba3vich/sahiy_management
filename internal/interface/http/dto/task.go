package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"

type CreateTaskRequest struct {
	ParentID    *int64 `json:"parent_id" example:"1"`
	SectionID   int64  `json:"section_id" binding:"required" example:"1"`
	Title       string `json:"title" binding:"required" example:"Implement login feature"`
	Description string `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int    `json:"priority" binding:"required" example:"1"`
	Deadline    int64  `json:"deadline" binding:"required" example:"1737363600"`
	Status      *int   `json:"status" example:"1"`
}

type UpdateTaskRequest struct {
	ParentID    *int64 `json:"parent_id" example:"1"`
	SectionID   int64  `json:"section_id" binding:"required" example:"1"`
	Title       string `json:"title" binding:"required" example:"Implement login feature"`
	Description string `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int    `json:"priority" binding:"required" example:"1"`
	Deadline    int64  `json:"deadline" binding:"required" example:"1737363600"`
	Status      int    `json:"status" binding:"required" example:"1"`
}

type TaskResponse struct {
	ID          int64  `json:"id" example:"1"`
	ParentID    *int64 `json:"parent_id,omitempty" example:"1"`
	SectionID   int64  `json:"section_id" example:"1"`
	Title       string `json:"title" example:"Implement login feature"`
	Description string `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int    `json:"priority" example:"1"`
	Deadline    int64  `json:"deadline" example:"1737363600"`
	Status      int    `json:"status" example:"1"`
	CreatedAt   int64  `json:"created_at" example:"1737277200"`
	UpdatedAt   int64  `json:"updated_at" example:"1737277200"`
}

func ToTaskResponse(task *domain.Task) *TaskResponse {
	resp := &TaskResponse{
		ID:          task.ID,
		SectionID:   task.SectionID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		Deadline:    task.Deadline,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	if task.ParentID.Valid {
		resp.ParentID = &task.ParentID.Int64
	}
	return resp
}

func ToTaskResponseList(tasks []*domain.Task) []*TaskResponse {
	responses := make([]*TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = ToTaskResponse(task)
	}
	return responses
}
