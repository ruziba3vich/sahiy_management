package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/section"

type CreateSectionRequest struct {
	DepartmentID int64  `json:"department_id" binding:"required" example:"1"`
	Name         string `json:"name" binding:"required" example:"Backend Development"`
}

type UpdateSectionRequest struct {
	DepartmentID int64  `json:"department_id" binding:"required" example:"1"`
	Name         string `json:"name" binding:"required" example:"Backend Development"`
}

type SectionResponse struct {
	ID           int64  `json:"id" example:"1"`
	DepartmentID int64  `json:"department_id" example:"1"`
	Name         string `json:"name" example:"Backend Development"`
	CreatedAt    int64  `json:"created_at" example:"1737277200"`
	UpdatedAt    int64  `json:"updated_at" example:"1737277200"`
}

func ToSectionResponse(sec *domain.Section) *SectionResponse {
	return &SectionResponse{
		ID:           sec.ID,
		DepartmentID: sec.DepartmentID,
		Name:         sec.Name,
		CreatedAt:    sec.CreatedAt,
		UpdatedAt:    sec.UpdatedAt,
	}
}

func ToSectionResponseList(sections []*domain.Section) []*SectionResponse {
	responses := make([]*SectionResponse, len(sections))
	for i, sec := range sections {
		responses[i] = ToSectionResponse(sec)
	}
	return responses
}
