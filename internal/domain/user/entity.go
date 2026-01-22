package user

import "time"

type User struct {
	ID           int64
	DepartmentID int64
	SectionID    int64
	ScheduleID   *int64
	Role         int
	Phone        string
	FullName     string
	JoinedAt     int64
	CreatedAt    int64
	UpdatedAt    int64
	TgChatID     int64
	PasswordHash string
}

func NewUser(departmentID, sectionID int64, scheduleID *int64, role int, phone, fullName, passwordHash string, tgChatID int64) *User {
	now := time.Now().Unix()
	return &User{
		DepartmentID: departmentID,
		SectionID:    sectionID,
		ScheduleID:   scheduleID,
		Role:         role,
		Phone:        phone,
		FullName:     fullName,
		JoinedAt:     now,
		CreatedAt:    now,
		UpdatedAt:    now,
		TgChatID:     tgChatID,
		PasswordHash: passwordHash,
	}
}

func (u *User) Update(departmentID, sectionID int64, scheduleID *int64, role int, phone, fullName string, joinedAt, tgChatID int64) {
	u.DepartmentID = departmentID
	u.SectionID = sectionID
	if scheduleID != nil {
		u.ScheduleID = scheduleID
	}
	u.Role = role
	u.Phone = phone
	u.FullName = fullName
	u.JoinedAt = joinedAt
	u.TgChatID = tgChatID
	u.UpdatedAt = time.Now().Unix()
}
