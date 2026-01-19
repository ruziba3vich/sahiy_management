package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/user"

type CreateUserRequest struct {
	DepartmentID int64  `json:"department_id" binding:"required" example:"1"`
	SectionID    int64  `json:"section_id" binding:"required" example:"1"`
	ScheduleID   int64  `json:"schedule_id" binding:"required" example:"1"`
	Role         int    `json:"role" binding:"required" example:"1"`
	Phone        string `json:"phone" binding:"required" example:"+998901234567"`
	FullName     string `json:"full_name" binding:"required" example:"John Doe"`
	TgChatID     int64  `json:"tg_chat_id" example:"123456789"`
	Password     string `json:"password" binding:"required" example:"securepassword"`
}

type UpdateUserRequest struct {
	DepartmentID int64  `json:"department_id" binding:"required" example:"1"`
	SectionID    int64  `json:"section_id" binding:"required" example:"1"`
	ScheduleID   int64  `json:"schedule_id" binding:"required" example:"1"`
	Role         int    `json:"role" binding:"required" example:"1"`
	Phone        string `json:"phone" binding:"required" example:"+998901234567"`
	FullName     string `json:"full_name" binding:"required" example:"John Doe"`
	TgChatID     int64  `json:"tg_chat_id" example:"123456789"`
}

type UserResponse struct {
	ID           int64  `json:"id" example:"1"`
	DepartmentID int64  `json:"department_id" example:"1"`
	SectionID    int64  `json:"section_id" example:"1"`
	ScheduleID   int64  `json:"schedule_id" example:"1"`
	Role         int    `json:"role" example:"1"`
	Phone        string `json:"phone" example:"+998901234567"`
	FullName     string `json:"full_name" example:"John Doe"`
	JoinedAt     int64  `json:"joined_at" example:"1737277200"`
	CreatedAt    int64  `json:"created_at" example:"1737277200"`
	UpdatedAt    int64  `json:"updated_at" example:"1737277200"`
	TgChatID     int64  `json:"tg_chat_id" example:"123456789"`
}

func ToUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:           user.ID,
		DepartmentID: user.DepartmentID,
		SectionID:    user.SectionID,
		ScheduleID:   user.ScheduleID,
		Role:         user.Role,
		Phone:        user.Phone,
		FullName:     user.FullName,
		JoinedAt:     user.JoinedAt,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		TgChatID:     user.TgChatID,
	}
}

func ToUserResponseList(users []*domain.User) []*UserResponse {
	responses := make([]*UserResponse, len(users))
	for i, user := range users {
		responses[i] = ToUserResponse(user)
	}
	return responses
}
