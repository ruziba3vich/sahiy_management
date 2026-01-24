package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appBranch "github.com/ruziba3vich/sahiy_management/internal/application/branch"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type BranchHandler struct {
	service *appBranch.Service
}

func NewBranchHandler(service *appBranch.Service) *BranchHandler {
	return &BranchHandler{service: service}
}

// Create godoc
// @Summary      Create a new branch
// @Description  Create a new branch with the provided data
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateBranchRequest  true  "Branch data"
// @Success      201      {object}  dto.BranchResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /branches [post] [deprecated]
func (h *BranchHandler) Create(c *gin.Context) {
	var req dto.CreateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	status := 1
	if req.Status != nil {
		status = *req.Status
	}

	branch, err := h.service.Create(c.Request.Context(), req.Name, req.Lat, req.Long, req.Radius, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToBranchResponse(branch))
}

// Update godoc
// @Summary      Update a branch
// @Description  Update an existing branch by ID
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "Branch ID"
// @Param        request  body      dto.UpdateBranchRequest  true  "Branch data"
// @Success      200      {object}  dto.BranchResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /branches/{id} [put] [deprecated]
func (h *BranchHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid branch ID"})
		return
	}

	var req dto.UpdateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	branch, err := h.service.Update(c.Request.Context(), id, req.Name, req.Lat, req.Long, req.Radius, req.Status)
	if err != nil {
		if errors.Is(err, postgres.ErrBranchNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "branch not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToBranchResponse(branch))
}

// Delete godoc
// @Summary      Delete a branch
// @Description  Delete a branch by ID
// @Tags         branches
// @Produce      json
// @Param        id  path  int  true  "Branch ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /branches/{id} [delete] [deprecated]
func (h *BranchHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid branch ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrBranchNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "branch not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a branch by ID
// @Description  Retrieve a single branch by its ID
// @Tags         branches
// @Produce      json
// @Param        id  path      int  true  "Branch ID"
// @Success      200  {object}  dto.BranchResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /branches/{id} [get] [deprecated]
func (h *BranchHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid branch ID"})
		return
	}

	branch, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrBranchNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "branch not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToBranchResponse(branch))
}

// GetAll godoc
// @Summary      Get all branches
// @Description  Retrieve a list of all branches
// @Tags         branches
// @Produce      json
// @Success      200  {array}   dto.BranchResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /branches [get] [deprecated]
func (h *BranchHandler) GetAll(c *gin.Context) {
	branches, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToBranchResponseList(branches))
}

func (h *BranchHandler) RegisterRoutes(r *gin.RouterGroup) {
	branches := r.Group("/branches")
	{
		branches.POST("", h.Create)
		branches.PUT("/:id", h.Update)
		branches.DELETE("/:id", h.Delete)
		branches.GET("/:id", h.GetByID)
		branches.GET("", h.GetAll)
	}
}
