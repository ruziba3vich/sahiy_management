package schedule

import (
	"encoding/json"
	"time"
)

type Schedule struct {
	ID        int64
	Name      string
	Timezone  int
	WeekDays  json.RawMessage
	Status    int
	CreatedAt int64
	UpdatedAt int64
}

func NewSchedule(name string, timezone int, weekDays json.RawMessage, status int) *Schedule {
	now := time.Now().Unix()
	return &Schedule{
		Name:      name,
		Timezone:  timezone,
		WeekDays:  weekDays,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s *Schedule) Update(name string, timezone int, weekDays json.RawMessage, status int) {
	s.Name = name
	s.Timezone = timezone
	s.WeekDays = weekDays
	s.Status = status
	s.UpdatedAt = time.Now().Unix()
}
