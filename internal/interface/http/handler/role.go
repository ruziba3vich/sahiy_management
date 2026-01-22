package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/domain/role"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type RoleHandler struct{}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{}
}

// GetAll godoc
// @Summary      Get all roles
// @Description  Retrieve a list of all available roles with their names and values
// @Tags         roles
// @Produce      json
// @Success      200  {array}   dto.RoleResponse
// @Router       /roles [get]
func (h *RoleHandler) GetAll(c *gin.Context) {
	roles := role.GetAllRoles()
	c.JSON(http.StatusOK, dto.ToRoleResponseList(roles))
}

func (h *RoleHandler) RegisterRoutes(r *gin.RouterGroup) {
	roles := r.Group("/roles")
	{
		roles.GET("", h.GetAll)
	}
}
