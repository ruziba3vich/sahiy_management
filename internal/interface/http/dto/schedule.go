package dto

import (
	"encoding/json"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/schedule"
)

type CreateScheduleRequest struct {
	Name     string          `json:"name" binding:"required" example:"Default"`
	Timezone int             `json:"timezone" binding:"required" example:"5"`
	WeekDays json.RawMessage `json:"week_days" binding:"required" example:"[1,2,3,4,5]"`
	Status   *int            `json:"status" example:"1"`
}

type UpdateScheduleRequest struct {
	Name     string          `json:"name" binding:"required" example:"Default"`
	Timezone int             `json:"timezone" binding:"required" example:"5"`
	WeekDays json.RawMessage `json:"week_days" binding:"required" example:"[1,2,3,4,5]"`
	Status   int             `json:"status" binding:"required" example:"1"`
}

type ScheduleResponse struct {
	ID        int64           `json:"id" example:"1"`
	Name      string          `json:"name" example:"Default"`
	Timezone  int             `json:"timezone" example:"5"`
	WeekDays  json.RawMessage `json:"week_days" example:"[1,2,3,4,5]"`
	Status    int             `json:"status" example:"1"`
	CreatedAt int64           `json:"created_at" example:"1737277200"`
	UpdatedAt int64           `json:"updated_at" example:"1737277200"`
}

func ToScheduleResponse(schedule *domain.Schedule) *ScheduleResponse {
	return &ScheduleResponse{
		ID:        schedule.ID,
		Name:      schedule.Name,
		Timezone:  schedule.Timezone,
		WeekDays:  schedule.WeekDays,
		Status:    schedule.Status,
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	}
}

func ToScheduleResponseList(schedules []*domain.Schedule) []*ScheduleResponse {
	responses := make([]*ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		responses[i] = ToScheduleResponse(schedule)
	}
	return responses
}
