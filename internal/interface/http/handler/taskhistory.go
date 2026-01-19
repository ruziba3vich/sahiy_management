package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appTH "github.com/ruziba3vich/sahiy_management/internal/application/taskhistory"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type TaskHistoryHandler struct {
	service *appTH.Service
}

func NewTaskHistoryHandler(service *appTH.Service) *TaskHistoryHandler {
	return &TaskHistoryHandler{service: service}
}

// Create godoc
// @Summary      Create a new task history
// @Description  Create a new task history entry
// @Tags         task-histories
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateTaskHistoryRequest  true  "Task history data"
// @Success      201      {object}  dto.TaskHistoryResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /task-histories [post]
func (h *TaskHistoryHandler) Create(c *gin.Context) {
	var req dto.CreateTaskHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	status := 1
	if req.Status != nil {
		status = *req.Status
	}

	th, err := h.service.Create(c.Request.Context(), req.TaskID, req.UserID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToTaskHistoryResponse(th))
}

// Update godoc
// @Summary      Update a task history
// @Description  Update an existing task history by ID
// @Tags         task-histories
// @Accept       json
// @Produce      json
// @Param        id       path      int                           true  "Task history ID"
// @Param        request  body      dto.UpdateTaskHistoryRequest  true  "Task history data"
// @Success      200      {object}  dto.TaskHistoryResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /task-histories/{id} [put]
func (h *TaskHistoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task history ID"})
		return
	}

	var req dto.UpdateTaskHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	th, err := h.service.Update(c.Request.Context(), id, req.TaskID, req.UserID, req.Status, req.FinishedAt)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskHistoryNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task history not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskHistoryResponse(th))
}

// Delete godoc
// @Summary      Delete a task history
// @Description  Delete a task history by ID
// @Tags         task-histories
// @Produce      json
// @Param        id  path  int  true  "Task history ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-histories/{id} [delete]
func (h *TaskHistoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task history ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskHistoryNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task history not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a task history by ID
// @Description  Retrieve a single task history by its ID
// @Tags         task-histories
// @Produce      json
// @Param        id  path      int  true  "Task history ID"
// @Success      200  {object}  dto.TaskHistoryResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-histories/{id} [get]
func (h *TaskHistoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task history ID"})
		return
	}

	th, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskHistoryNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task history not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskHistoryResponse(th))
}

// GetAll godoc
// @Summary      Get all task histories
// @Description  Retrieve a list of all task histories
// @Tags         task-histories
// @Produce      json
// @Success      200  {array}   dto.TaskHistoryResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-histories [get]
func (h *TaskHistoryHandler) GetAll(c *gin.Context) {
	histories, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskHistoryResponseList(histories))
}

func (h *TaskHistoryHandler) RegisterRoutes(r *gin.RouterGroup) {
	taskHistories := r.Group("/task-histories")
	{
		taskHistories.POST("", h.Create)
		taskHistories.PUT("/:id", h.Update)
		taskHistories.DELETE("/:id", h.Delete)
		taskHistories.GET("/:id", h.GetByID)
		taskHistories.GET("", h.GetAll)
	}
}
