package dto

import domain "github.com/ruziba3vich/sahiy_management/internal/domain/branch"

type CreateBranchRequest struct {
	Name   string  `json:"name" binding:"required" example:"Main Branch"`
	Lat    float64 `json:"lat" binding:"required" example:"41.311081"`
	Long   float64 `json:"long" binding:"required" example:"69.240562"`
	Radius int     `json:"radius" binding:"required" example:"150"`
	Status *int    `json:"status" example:"1"`
}

type UpdateBranchRequest struct {
	Name   string  `json:"name" binding:"required" example:"Main Branch"`
	Lat    float64 `json:"lat" binding:"required" example:"41.311081"`
	Long   float64 `json:"long" binding:"required" example:"69.240562"`
	Radius int     `json:"radius" binding:"required" example:"150"`
	Status int     `json:"status" binding:"required" example:"1"`
}

type BranchResponse struct {
	ID        int64   `json:"id" example:"1"`
	Name      string  `json:"name" example:"Main Branch"`
	Lat       float64 `json:"lat" example:"41.311081"`
	Long      float64 `json:"long" example:"69.240562"`
	Radius    int     `json:"radius" example:"150"`
	Status    int     `json:"status" example:"1"`
	CreatedAt int64   `json:"created_at" example:"1737277200"`
	UpdatedAt int64   `json:"updated_at" example:"1737277200"`
}

func ToBranchResponse(branch *domain.Branch) *BranchResponse {
	return &BranchResponse{
		ID:        branch.ID,
		Name:      branch.Name,
		Lat:       branch.Lat,
		Long:      branch.Long,
		Radius:    branch.Radius,
		Status:    branch.Status,
		CreatedAt: branch.CreatedAt,
		UpdatedAt: branch.UpdatedAt,
	}
}

func ToBranchResponseList(branches []*domain.Branch) []*BranchResponse {
	responses := make([]*BranchResponse, len(branches))
	for i, branch := range branches {
		responses[i] = ToBranchResponse(branch)
	}
	return responses
}
