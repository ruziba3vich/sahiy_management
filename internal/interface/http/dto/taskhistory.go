package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"

type CreateTaskHistoryRequest struct {
	TaskID int64 `json:"task_id" binding:"required" example:"1"`
	UserID int64 `json:"user_id" binding:"required" example:"1"`
	Status *int  `json:"status" example:"1"`
}

type UpdateTaskHistoryRequest struct {
	TaskID     int64  `json:"task_id" binding:"required" example:"1"`
	UserID     int64  `json:"user_id" binding:"required" example:"1"`
	Status     int    `json:"status" binding:"required" example:"1"`
	FinishedAt *int64 `json:"finished_at" example:"1737363600"`
}

type TaskHistoryResponse struct {
	ID         int64  `json:"id" example:"1"`
	TaskID     int64  `json:"task_id" example:"1"`
	UserID     int64  `json:"user_id" example:"1"`
	Status     int    `json:"status" example:"1"`
	StartedAt  int64  `json:"started_at" example:"1737277200"`
	FinishedAt *int64 `json:"finished_at,omitempty" example:"1737363600"`
}

func ToTaskHistoryResponse(th *domain.TaskHistory) *TaskHistoryResponse {
	resp := &TaskHistoryResponse{
		ID:        th.ID,
		TaskID:    th.TaskID,
		UserID:    th.UserID,
		Status:    th.Status,
		StartedAt: th.StartedAt,
	}
	if th.FinishedAt.Valid {
		resp.FinishedAt = &th.FinishedAt.Int64
	}
	return resp
}

func ToTaskHistoryResponseList(histories []*domain.TaskHistory) []*TaskHistoryResponse {
	responses := make([]*TaskHistoryResponse, len(histories))
	for i, th := range histories {
		responses[i] = ToTaskHistoryResponse(th)
	}
	return responses
}

type TaskHistoryListResponse struct {
	Data       []*TaskHistoryResponse `json:"data"`
	TotalCount int64                  `json:"total_count" example:"100"`
	Page       int                    `json:"page" example:"1"`
	PageSize   int                    `json:"page_size" example:"20"`
	TotalPages int                    `json:"total_pages" example:"5"`
}

func ToTaskHistoryListResponse(result *domain.TaskHistoryListResult, page, pageSize int) *TaskHistoryListResponse {
	totalPages := int(result.TotalCount) / pageSize
	if int(result.TotalCount)%pageSize > 0 {
		totalPages++
	}

	return &TaskHistoryListResponse{
		Data:       ToTaskHistoryResponseList(result.Histories),
		TotalCount: result.TotalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}
