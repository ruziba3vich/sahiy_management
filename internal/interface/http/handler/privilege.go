package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/application/privilege"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type PrivilegeHandler struct {
	service *privilege.Service
}

func NewPrivilegeHandler(service *privilege.Service) *PrivilegeHandler {
	return &PrivilegeHandler{service: service}
}

func (h *PrivilegeHandler) RegisterRoutes(rg *gin.RouterGroup) {
	privileges := rg.Group("/privileges")
	{
		privileges.POST("", h.Create)
		privileges.GET("", h.GetAll)
		privileges.GET("/:id", h.GetByID)
		privileges.PUT("/:id", h.Update)
		privileges.DELETE("/:id", h.Delete)
	}
}

func (h *PrivilegeHandler) RegisterUserPrivilegeRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/:id/privileges", h.GetUserPrivileges)
		users.POST("/:id/privileges", h.AssignPrivilege)
		users.DELETE("/:id/privileges/:privilegeId", h.RevokePrivilege)
	}
}

// Create godoc
// @Summary      Create privilege
// @Description  Create a new privilege
// @Tags         privileges
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreatePrivilegeRequest true "Privilege data"
// @Success      201 {object} dto.PrivilegeResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /privileges [post]
func (h *PrivilegeHandler) Create(c *gin.Context) {
	var req dto.CreatePrivilegeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.service.Create(c.Request.Context(), req.Name, req.Resource, req.Action, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToPrivilegeResponse(p))
}

// GetAll godoc
// @Summary      Get all privileges
// @Description  Get all privileges
// @Tags         privileges
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} dto.PrivilegeResponse
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /privileges [get]
func (h *PrivilegeHandler) GetAll(c *gin.Context) {
	privileges, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToPrivilegeResponseList(privileges))
}

// GetByID godoc
// @Summary      Get privilege by ID
// @Description  Get a privilege by its ID
// @Tags         privileges
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Privilege ID"
// @Success      200 {object} dto.PrivilegeResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /privileges/{id} [get]
func (h *PrivilegeHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	p, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "privilege not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToPrivilegeResponse(p))
}

// Update godoc
// @Summary      Update privilege
// @Description  Update an existing privilege
// @Tags         privileges
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Privilege ID"
// @Param        request body dto.UpdatePrivilegeRequest true "Privilege data"
// @Success      200 {object} dto.PrivilegeResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /privileges/{id} [put]
func (h *PrivilegeHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdatePrivilegeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.service.Update(c.Request.Context(), id, req.Name, req.Resource, req.Action, req.Description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "privilege not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToPrivilegeResponse(p))
}

// Delete godoc
// @Summary      Delete privilege
// @Description  Delete a privilege by its ID
// @Tags         privileges
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Privilege ID"
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /privileges/{id} [delete]
func (h *PrivilegeHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "privilege not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUserPrivileges godoc
// @Summary      Get user privileges
// @Description  Get all privileges assigned to a user
// @Tags         user-privileges
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "User ID"
// @Success      200 {object} dto.UserPrivilegeResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /users/{id}/privileges [get]
func (h *PrivilegeHandler) GetUserPrivileges(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	privileges, err := h.service.GetUserPrivileges(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserPrivilegeResponse(userID, privileges))
}

// AssignPrivilege godoc
// @Summary      Assign privilege to user
// @Description  Assign a privilege to a user
// @Tags         user-privileges
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "User ID"
// @Param        request body dto.AssignPrivilegeRequest true "Privilege assignment data"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /users/{id}/privileges [post]
func (h *PrivilegeHandler) AssignPrivilege(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req dto.AssignPrivilegeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grantedByID, _ := c.Get("userID")
	var grantedBy *int64
	if id, ok := grantedByID.(int64); ok {
		grantedBy = &id
	}

	_, err = h.service.AssignToUser(c.Request.Context(), userID, req.PrivilegeID, grantedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "privilege assigned successfully"})
}

// RevokePrivilege godoc
// @Summary      Revoke privilege from user
// @Description  Revoke a privilege from a user
// @Tags         user-privileges
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "User ID"
// @Param        privilegeId path int true "Privilege ID"
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /users/{id}/privileges/{privilegeId} [delete]
func (h *PrivilegeHandler) RevokePrivilege(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	privilegeID, err := strconv.ParseInt(c.Param("privilegeId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid privilege id"})
		return
	}

	if err := h.service.RevokeFromUser(c.Request.Context(), userID, privilegeID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user privilege not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
