package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appDept "github.com/ruziba3vich/sahiy_management/internal/application/department"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type DepartmentHandler struct {
	service *appDept.Service
}

func NewDepartmentHandler(service *appDept.Service) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

// Create godoc
// @Summary      Create a new department
// @Description  Create a new department with the provided name and status
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateDepartmentRequest  true  "Department data"
// @Success      201      {object}  dto.DepartmentResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /departments [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	dept, err := h.service.Create(c.Request.Context(), req.Name, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToDepartmentResponse(dept))
}

// Update godoc
// @Summary      Update a department
// @Description  Update an existing department by ID
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Department ID"
// @Param        request  body      dto.UpdateDepartmentRequest  true  "Department data"
// @Success      200      {object}  dto.DepartmentResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /departments/{id} [put]
func (h *DepartmentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid department ID"})
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	dept, err := h.service.Update(c.Request.Context(), id, req.Name, req.Status)
	if err != nil {
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "department not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDepartmentResponse(dept))
}

// Delete godoc
// @Summary      Delete a department
// @Description  Delete a department by ID
// @Tags         departments
// @Produce      json
// @Param        id  path  int  true  "Department ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /departments/{id} [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid department ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "department not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a department by ID
// @Description  Retrieve a single department by its ID
// @Tags         departments
// @Produce      json
// @Param        id  path      int  true  "Department ID"
// @Success      200  {object}  dto.DepartmentResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /departments/{id} [get]
func (h *DepartmentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid department ID"})
		return
	}

	dept, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "department not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDepartmentResponse(dept))
}

// GetAll godoc
// @Summary      Get all departments
// @Description  Retrieve a list of all departments
// @Tags         departments
// @Produce      json
// @Success      200  {array}   dto.DepartmentResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /departments [get]
func (h *DepartmentHandler) GetAll(c *gin.Context) {
	depts, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDepartmentResponseList(depts))
}

func (h *DepartmentHandler) RegisterRoutes(r *gin.RouterGroup) {
	departments := r.Group("/departments")
	{
		departments.POST("", h.Create)
		departments.PUT("/:id", h.Update)
		departments.DELETE("/:id", h.Delete)
		departments.GET("/:id", h.GetByID)
		departments.GET("", h.GetAll)
	}
}
