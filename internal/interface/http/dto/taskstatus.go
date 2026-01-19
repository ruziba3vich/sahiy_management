package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"

type CreateTaskStatusRequest struct {
	Name string `json:"name" binding:"required" example:"In Progress"`
	Type int    `json:"type" binding:"required" example:"1"`
}

type UpdateTaskStatusRequest struct {
	Name string `json:"name" binding:"required" example:"In Progress"`
	Type int    `json:"type" binding:"required" example:"1"`
}

type TaskStatusResponse struct {
	ID   int64  `json:"id" example:"1"`
	Name string `json:"name" example:"In Progress"`
	Type int    `json:"type" example:"1"`
}

func ToTaskStatusResponse(ts *domain.TaskStatus) *TaskStatusResponse {
	return &TaskStatusResponse{
		ID:   ts.ID,
		Name: ts.Name,
		Type: ts.Type,
	}
}

func ToTaskStatusResponseList(statuses []*domain.TaskStatus) []*TaskStatusResponse {
	responses := make([]*TaskStatusResponse, len(statuses))
	for i, ts := range statuses {
		responses[i] = ToTaskStatusResponse(ts)
	}
	return responses
}
