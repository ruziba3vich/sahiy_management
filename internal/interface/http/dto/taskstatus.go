package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"

type CreateTaskStatusRequest struct {
	Name      string `json:"name" binding:"required" example:"In Progress"`
	SectionID int64  `json:"section_id" binding:"required" example:"1"`
	Color     string `json:"color" example:"#ffffff"`
}

type UpdateTaskStatusRequest struct {
	Name      string `json:"name" binding:"required" example:"In Progress"`
	SectionID int64  `json:"section_id" binding:"required" example:"1"`
	Color     string `json:"color" example:"#ffffff"`
}

type TaskStatusResponse struct {
	ID        int64  `json:"id" example:"1"`
	Name      string `json:"name" example:"In Progress"`
	SectionID int64  `json:"section_id" example:"1"`
	Color     string `json:"color" example:"#ffffff"`
}

func ToTaskStatusResponse(ts *domain.TaskStatus) *TaskStatusResponse {
	return &TaskStatusResponse{
		ID:        ts.ID,
		Name:      ts.Name,
		SectionID: ts.SectionID,
		Color:     ts.Color,
	}
}

func ToTaskStatusResponseList(statuses []*domain.TaskStatus) []*TaskStatusResponse {
	responses := make([]*TaskStatusResponse, len(statuses))
	for i, ts := range statuses {
		responses[i] = ToTaskStatusResponse(ts)
	}
	return responses
}
