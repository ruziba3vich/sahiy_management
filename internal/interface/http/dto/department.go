package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/department"

type CreateDepartmentRequest struct {
	Name   string `json:"name" binding:"required" example:"Engineering"`
	Status *int   `json:"status" example:"1"`
}

type UpdateDepartmentRequest struct {
	Name   string `json:"name" binding:"required" example:"Engineering"`
	Status int    `json:"status" binding:"required" example:"1"`
}

type DepartmentResponse struct {
	ID        int64  `json:"id" example:"1"`
	Name      string `json:"name" example:"Engineering"`
	Status    int    `json:"status" example:"1"`
	CreatedAt int64  `json:"created_at" example:"1737277200"`
	UpdatedAt int64  `json:"updated_at" example:"1737277200"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"department not found"`
}

func ToDepartmentResponse(dept *domain.Department) *DepartmentResponse {
	return &DepartmentResponse{
		ID:        dept.ID,
		Name:      dept.Name,
		Status:    dept.Status,
		CreatedAt: dept.CreatedAt,
		UpdatedAt: dept.UpdatedAt,
	}
}

func ToDepartmentResponseList(depts []*domain.Department) []*DepartmentResponse {
	responses := make([]*DepartmentResponse, len(depts))
	for i, dept := range depts {
		responses[i] = ToDepartmentResponse(dept)
	}
	return responses
}
