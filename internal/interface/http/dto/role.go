package dto

import "github.com/ruziba3vich/sahiy_management/internal/domain/role"

type RoleResponse struct {
	Name  string `json:"name" example:"Super Admin"`
	Value int    `json:"value" example:"99"`
}

func ToRoleResponse(r role.Role) RoleResponse {
	return RoleResponse{
		Name:  r.Name,
		Value: r.Value,
	}
}

func ToRoleResponseList(roles []role.Role) []RoleResponse {
	responses := make([]RoleResponse, len(roles))
	for i, r := range roles {
		responses[i] = ToRoleResponse(r)
	}
	return responses
}
