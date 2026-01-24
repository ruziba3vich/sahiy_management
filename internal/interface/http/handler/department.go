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

	status := 1
	if req.Status != nil {
		status = *req.Status
	}

	dept, err := h.service.Create(c.Request.Context(), req.Name, status)
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

// AttachBranch godoc
// @Summary      Attach a branch to a department
// @Description  Create a link between a branch and a department
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AttachBranchRequest  true  "Attachment data"
// @Success      201      {object}  dto.DepartmentBranchResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /departments/branches [post]
func (h *DepartmentHandler) AttachBranch(c *gin.Context) {
	var req dto.AttachBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	db, err := h.service.AttachBranch(c.Request.Context(), req.BranchID, req.DepartmentID)
	if err != nil {
		if errors.Is(err, postgres.ErrDuplicateDepartmentBranch) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToDepartmentBranchResponse(db))
}

// UpdateDepartmentBranchStatus godoc
// @Summary      Update department branch status
// @Description  Update the status of a branch-department link
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        id       path      int                                       true  "Branch-Department Link ID"
// @Param        request  body      dto.UpdateDepartmentBranchStatusRequest  true  "Status data"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /departments/branches/{id}/status [put]
func (h *DepartmentHandler) UpdateDepartmentBranchStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid link ID"})
		return
	}

	var req dto.UpdateDepartmentBranchStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	err = h.service.UpdateDepartmentBranchStatus(c.Request.Context(), id, *req.Status)
	if err != nil {
		if errors.Is(err, postgres.ErrDepartmentBranchNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "department branch link not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}

// GetBranchesByDepartmentID godoc
// @Summary      Get branches for a department
// @Description  Retrieve all branch links for a given department
// @Tags         departments
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Success      200  {array}   dto.DepartmentBranchResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /departments/{id}/branches [get]
func (h *DepartmentHandler) GetBranchesByDepartmentID(c *gin.Context) {
	departmentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid department ID"})
		return
	}

	branches, err := h.service.GetBranchesByDepartmentID(c.Request.Context(), departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDepartmentBranchResponseList(branches))
}

func (h *DepartmentHandler) RegisterRoutes(r *gin.RouterGroup) {
	departments := r.Group("/departments")
	{
		departments.POST("/branches", h.AttachBranch)
		departments.PUT("/branches/:id/status", h.UpdateDepartmentBranchStatus)
		departments.GET("/:id/branches", h.GetBranchesByDepartmentID)

		departments.POST("", h.Create)
		departments.PUT("/:id", h.Update)
		departments.DELETE("/:id", h.Delete)
		departments.GET("/:id", h.GetByID)
		departments.GET("", h.GetAll)
	}
}
