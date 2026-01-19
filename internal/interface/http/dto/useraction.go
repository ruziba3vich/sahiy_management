package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/useraction"

type CreateUserActionRequest struct {
	UserID        int64  `json:"user_id" binding:"required" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" binding:"required" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    int    `json:"come_status" binding:"required" example:"1"`
	OutStatus     *int   `json:"out_status" example:"1"`
	StartedAt     int64  `json:"started_at" binding:"required" example:"1737277200"`
	FinishedAt    *int64 `json:"finished_at" example:"1737280800"`
}

type UpdateUserActionRequest struct {
	UserID        int64  `json:"user_id" binding:"required" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" binding:"required" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    int    `json:"come_status" binding:"required" example:"1"`
	OutStatus     *int   `json:"out_status" example:"1"`
	StartedAt     int64  `json:"started_at" binding:"required" example:"1737277200"`
	FinishedAt    *int64 `json:"finished_at" example:"1737280800"`
}

type UserActionResponse struct {
	ID            int64  `json:"id" example:"1"`
	UserID        int64  `json:"user_id" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    int    `json:"come_status" example:"1"`
	OutStatus     *int   `json:"out_status" example:"1"`
	StartedAt     int64  `json:"started_at" example:"1737277200"`
	FinishedAt    *int64 `json:"finished_at" example:"1737280800"`
}

func ToUserActionResponse(action *domain.UserAction) *UserActionResponse {
	return &UserActionResponse{
		ID:            action.ID,
		UserID:        action.UserID,
		VisitBranchID: action.VisitBranchID,
		LeaveBranchID: action.LeaveBranchID,
		ComeStatus:    action.ComeStatus,
		OutStatus:     action.OutStatus,
		StartedAt:     action.StartedAt,
		FinishedAt:    action.FinishedAt,
	}
}

func ToUserActionResponseList(actions []*domain.UserAction) []*UserActionResponse {
	responses := make([]*UserActionResponse, len(actions))
	for i, action := range actions {
		responses[i] = ToUserActionResponse(action)
	}
	return responses
}
