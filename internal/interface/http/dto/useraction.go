package dto

import (
	"fmt"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/useraction"
)

type CreateUserActionRequest struct {
	UserID        int64  `json:"user_id" binding:"required" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" binding:"required" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    *int   `json:"come_status" binding:"required" example:"1"`
	OutStatus     *int   `json:"out_status" example:"1"`
	StartedAt     int64  `json:"started_at" binding:"required" example:"1737277200"`
	FinishedAt    *int64 `json:"finished_at" example:"1737280800"`
}

type AcceptUserActionRequest struct {
	UserID int64   `json:"user_id" binding:"required" example:"1"`
	Lat    float64 `json:"lat" binding:"required" example:"41.311081"`
	Long   float64 `json:"long" binding:"required" example:"69.240562"`
}

type UpdateUserActionRequest struct {
	UserID        int64  `json:"user_id" binding:"required" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" binding:"required" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    *int   `json:"come_status" binding:"required" example:"1"`
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

type AttendanceStatusResponse struct {
	ID    int    `json:"id" example:"1"`
	Label string `json:"label" example:"O'z vaqtida"`
}

type UserAttendanceResponse struct {
	UserID          int64              `json:"user_id" example:"1"`
	UserFullName    string             `json:"user_full_name" example:"John Doe"`
	UserPhone       string             `json:"user_phone" example:"+998901234567"`
	TotalWorkedTime string             `json:"total_worked_time" example:"16:30"`
	Actions         []AttendanceAction `json:"actions"`
}

type AttendanceAction struct {
	ID            int64  `json:"id" example:"1"`
	VisitBranchID int64  `json:"visit_branch_id" example:"1"`
	LeaveBranchID *int64 `json:"leave_branch_id" example:"2"`
	ComeStatus    int    `json:"come_status" example:"1"`
	OutStatus     *int   `json:"out_status" example:"1"`
	StartedAt     int64  `json:"started_at" example:"1737277200"`
	FinishedAt    *int64 `json:"finished_at" example:"1737280800"`
	WorkedTime    string `json:"worked_time" example:"08:30"`
}

func ToUserAttendanceResponseList(records []*domain.AttendanceRecord) []*UserAttendanceResponse {
	userMap := make(map[int64]*UserAttendanceResponse)
	userOrder := []int64{}

	for _, record := range records {
		if _, ok := userMap[record.UserID]; !ok {
			userMap[record.UserID] = &UserAttendanceResponse{
				UserID:       record.UserID,
				UserFullName: record.UserFullName,
				UserPhone:    record.UserPhone,
				Actions:      []AttendanceAction{},
			}
			userOrder = append(userOrder, record.UserID)
		}

		action := AttendanceAction{
			ID:            record.ID,
			VisitBranchID: record.VisitBranchID,
			LeaveBranchID: record.LeaveBranchID,
			ComeStatus:    record.ComeStatus,
			OutStatus:     record.OutStatus,
			StartedAt:     record.StartedAt,
			FinishedAt:    record.FinishedAt,
		}

		if record.FinishedAt != nil {
			duration := *record.FinishedAt - record.StartedAt
			hours := duration / 3600
			minutes := (duration % 3600) / 60
			action.WorkedTime = fmt.Sprintf("%02d:%02d", hours, minutes)
		}

		userMap[record.UserID].Actions = append(userMap[record.UserID].Actions, action)
	}

	responses := make([]*UserAttendanceResponse, 0, len(userOrder))
	for _, userID := range userOrder {
		resp := userMap[userID]
		var totalDuration int64
		for _, action := range resp.Actions {
			if action.FinishedAt != nil {
				totalDuration += (*action.FinishedAt - action.StartedAt)
			}
		}
		hours := totalDuration / 3600
		minutes := (totalDuration % 3600) / 60
		resp.TotalWorkedTime = fmt.Sprintf("%02d:%02d", hours, minutes)
		responses = append(responses, resp)
	}

	return responses
}
