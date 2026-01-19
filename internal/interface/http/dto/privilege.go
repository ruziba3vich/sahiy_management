package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/privilege"

type CreatePrivilegeRequest struct {
	Name        string `json:"name" binding:"required" example:"departments:create"`
	Resource    string `json:"resource" binding:"required" example:"departments"`
	Action      string `json:"action" binding:"required" example:"create"`
	Description string `json:"description" example:"Create new departments"`
}

type UpdatePrivilegeRequest struct {
	Name        string `json:"name" binding:"required" example:"departments:create"`
	Resource    string `json:"resource" binding:"required" example:"departments"`
	Action      string `json:"action" binding:"required" example:"create"`
	Description string `json:"description" example:"Create new departments"`
}

type PrivilegeResponse struct {
	ID          int64  `json:"id" example:"1"`
	Name        string `json:"name" example:"departments:create"`
	Resource    string `json:"resource" example:"departments"`
	Action      string `json:"action" example:"create"`
	Description string `json:"description" example:"Create new departments"`
	CreatedAt   int64  `json:"created_at" example:"1737277200"`
	UpdatedAt   int64  `json:"updated_at" example:"1737277200"`
}

type AssignPrivilegeRequest struct {
	PrivilegeID int64 `json:"privilege_id" binding:"required" example:"1"`
}

type UserPrivilegeResponse struct {
	UserID     int64               `json:"user_id" example:"1"`
	Privileges []PrivilegeResponse `json:"privileges"`
}

func ToPrivilegeResponse(p *domain.Privilege) *PrivilegeResponse {
	return &PrivilegeResponse{
		ID:          p.ID,
		Name:        p.Name,
		Resource:    p.Resource,
		Action:      p.Action,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ToPrivilegeResponseList(privileges []*domain.Privilege) []*PrivilegeResponse {
	responses := make([]*PrivilegeResponse, len(privileges))
	for i, p := range privileges {
		responses[i] = ToPrivilegeResponse(p)
	}
	return responses
}

func ToUserPrivilegeResponse(userID int64, privileges []*domain.Privilege) *UserPrivilegeResponse {
	privResponses := make([]PrivilegeResponse, len(privileges))
	for i, p := range privileges {
		privResponses[i] = *ToPrivilegeResponse(p)
	}
	return &UserPrivilegeResponse{
		UserID:     userID,
		Privileges: privResponses,
	}
}
