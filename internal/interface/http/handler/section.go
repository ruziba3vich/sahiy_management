package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appSec "github.com/ruziba3vich/sahiy_management/internal/application/section"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type SectionHandler struct {
	service *appSec.Service
}

func NewSectionHandler(service *appSec.Service) *SectionHandler {
	return &SectionHandler{service: service}
}

// Create godoc
// @Summary      Create a new section
// @Description  Create a new section with the provided department_id and name
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateSectionRequest  true  "Section data"
// @Success      201      {object}  dto.SectionResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /sections [post]
func (h *SectionHandler) Create(c *gin.Context) {
	var req dto.CreateSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	sec, err := h.service.Create(c.Request.Context(), req.DepartmentID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToSectionResponse(sec))
}

// Update godoc
// @Summary      Update a section
// @Description  Update an existing section by ID
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        id       path      int                       true  "Section ID"
// @Param        request  body      dto.UpdateSectionRequest  true  "Section data"
// @Success      200      {object}  dto.SectionResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /sections/{id} [put]
func (h *SectionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid section ID"})
		return
	}

	var req dto.UpdateSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	sec, err := h.service.Update(c.Request.Context(), id, req.DepartmentID, req.Name)
	if err != nil {
		if errors.Is(err, postgres.ErrSectionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "section not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToSectionResponse(sec))
}

// Delete godoc
// @Summary      Delete a section
// @Description  Delete a section by ID
// @Tags         sections
// @Produce      json
// @Param        id  path  int  true  "Section ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /sections/{id} [delete]
func (h *SectionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid section ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrSectionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "section not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a section by ID
// @Description  Retrieve a single section by its ID
// @Tags         sections
// @Produce      json
// @Param        id  path      int  true  "Section ID"
// @Success      200  {object}  dto.SectionResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /sections/{id} [get]
func (h *SectionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid section ID"})
		return
	}

	sec, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrSectionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "section not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToSectionResponse(sec))
}

// GetAll godoc
// @Summary      Get all sections
// @Description  Retrieve a list of all sections
// @Tags         sections
// @Produce      json
// @Success      200  {array}   dto.SectionResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /sections [get]
func (h *SectionHandler) GetAll(c *gin.Context) {
	sections, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToSectionResponseList(sections))
}

func (h *SectionHandler) RegisterRoutes(r *gin.RouterGroup) {
	sections := r.Group("/sections")
	{
		sections.POST("", h.Create)
		sections.PUT("/:id", h.Update)
		sections.DELETE("/:id", h.Delete)
		sections.GET("/:id", h.GetByID)
		sections.GET("", h.GetAll)
	}
}
