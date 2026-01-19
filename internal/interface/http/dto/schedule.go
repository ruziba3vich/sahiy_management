package dto

import (
	"encoding/json"

	domain "github.com/ruziba3vich/sahiy_management/internal/domain/schedule"
)

type CreateScheduleRequest struct {
	Name     string                 `json:"name" binding:"required" example:"Default"`
	Timezone int                    `json:"timezone" binding:"required" example:"5"`
	WeekDays map[string]ScheduleDay `json:"week_days" binding:"required"`
	Status   *int                   `json:"status" example:"1"`
}

type UpdateScheduleRequest struct {
	Name     string                 `json:"name" binding:"required" example:"Default"`
	Timezone int                    `json:"timezone" binding:"required" example:"5"`
	WeekDays map[string]ScheduleDay `json:"week_days" binding:"required"`
	Status   int                    `json:"status" binding:"required" example:"1"`
}

type ScheduleResponse struct {
	ID        int64                  `json:"id" example:"1"`
	Name      string                 `json:"name" example:"Default"`
	Timezone  int                    `json:"timezone" example:"5"`
	WeekDays  map[string]ScheduleDay `json:"week_days"`
	Status    int                    `json:"status" example:"1"`
	CreatedAt int64                  `json:"created_at" example:"1737277200"`
	UpdatedAt int64                  `json:"updated_at" example:"1737277200"`
}

type ScheduleDay struct {
	StartAt  string `json:"start_at" example:"09:00"`
	FinishAt string `json:"finish_at" example:"18:00"`
}

func ToScheduleResponse(schedule *domain.Schedule) *ScheduleResponse {
	var weekDays map[string]ScheduleDay
	_ = json.Unmarshal(schedule.WeekDays, &weekDays)
	return &ScheduleResponse{
		ID:        schedule.ID,
		Name:      schedule.Name,
		Timezone:  schedule.Timezone,
		WeekDays:  weekDays,
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
