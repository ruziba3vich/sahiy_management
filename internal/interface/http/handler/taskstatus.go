package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appTS "github.com/ruziba3vich/sahiy_management/internal/application/taskstatus"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type TaskStatusHandler struct {
	service *appTS.Service
}

func NewTaskStatusHandler(service *appTS.Service) *TaskStatusHandler {
	return &TaskStatusHandler{service: service}
}

// Create godoc
// @Summary      Create a new task status
// @Description  Create a new task status with the provided name and type
// @Tags         task-statuses
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateTaskStatusRequest  true  "Task status data"
// @Success      201      {object}  dto.TaskStatusResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /task-statuses [post]
func (h *TaskStatusHandler) Create(c *gin.Context) {
	var req dto.CreateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	ts, err := h.service.Create(c.Request.Context(), req.Name, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToTaskStatusResponse(ts))
}

// Update godoc
// @Summary      Update a task status
// @Description  Update an existing task status by ID
// @Tags         task-statuses
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Task status ID"
// @Param        request  body      dto.UpdateTaskStatusRequest  true  "Task status data"
// @Success      200      {object}  dto.TaskStatusResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /task-statuses/{id} [put]
func (h *TaskStatusHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task status ID"})
		return
	}

	var req dto.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	ts, err := h.service.Update(c.Request.Context(), id, req.Name, req.Type)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskStatusNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task status not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskStatusResponse(ts))
}

// Delete godoc
// @Summary      Delete a task status
// @Description  Delete a task status by ID
// @Tags         task-statuses
// @Produce      json
// @Param        id  path  int  true  "Task status ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-statuses/{id} [delete]
func (h *TaskStatusHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task status ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskStatusNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task status not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a task status by ID
// @Description  Retrieve a single task status by its ID
// @Tags         task-statuses
// @Produce      json
// @Param        id  path      int  true  "Task status ID"
// @Success      200  {object}  dto.TaskStatusResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-statuses/{id} [get]
func (h *TaskStatusHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task status ID"})
		return
	}

	ts, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskStatusNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task status not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskStatusResponse(ts))
}

// GetAll godoc
// @Summary      Get all task statuses
// @Description  Retrieve a list of all task statuses
// @Tags         task-statuses
// @Produce      json
// @Success      200  {array}   dto.TaskStatusResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-statuses [get]
func (h *TaskStatusHandler) GetAll(c *gin.Context) {
	statuses, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskStatusResponseList(statuses))
}

func (h *TaskStatusHandler) RegisterRoutes(r *gin.RouterGroup) {
	taskStatuses := r.Group("/task-statuses")
	{
		taskStatuses.POST("", h.Create)
		taskStatuses.PUT("/:id", h.Update)
		taskStatuses.DELETE("/:id", h.Delete)
		taskStatuses.GET("/:id", h.GetByID)
		taskStatuses.GET("", h.GetAll)
	}
}
