package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appTH "github.com/ruziba3vich/sahiy_management/internal/application/taskhistory"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/taskhistory"
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
// @Summary      Get all task histories with filtering and pagination
// @Description  Retrieve a list of task histories with optional filters (user_id, task_id, status) and pagination
// @Tags         task-histories
// @Produce      json
// @Security     BearerAuth
// @Param        user_id    query     int  false  "Filter by user ID"
// @Param        task_id    query     int  false  "Filter by task ID"
// @Param        status     query     int  false  "Filter by status"
// @Param        page       query     int  false  "Page number (default: 1)"
// @Param        page_size  query     int  false  "Page size (default: 20, max: 100)"
// @Success      200  {object}  dto.TaskHistoryListResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-histories [get]
func (h *TaskHistoryHandler) GetAll(c *gin.Context) {
	filter := &domain.TaskHistoryFilter{
		Page:     1,
		PageSize: 20,
	}

	if userID := c.Query("user_id"); userID != "" {
		if id, err := strconv.ParseInt(userID, 10, 64); err == nil {
			filter.UserID = &id
		}
	}

	if taskID := c.Query("task_id"); taskID != "" {
		if id, err := strconv.ParseInt(taskID, 10, 64); err == nil {
			filter.TaskID = &id
		}
	}

	if status := c.Query("status"); status != "" {
		if s, err := strconv.Atoi(status); err == nil {
			filter.Status = &s
		}
	}

	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			filter.Page = p
		}
	}

	if pageSize := c.Query("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
			if ps > 100 {
				ps = 100
			}
			filter.PageSize = ps
		}
	}

	result, err := h.service.GetAllWithFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskHistoryListResponse(result, filter.Page, filter.PageSize))
}

// GetMyHistory godoc
// @Summary      Get current user's task histories with filtering and pagination
// @Description  Retrieve task histories for the authenticated user with optional filters (task_id, status) and pagination
// @Tags         task-histories
// @Produce      json
// @Security     BearerAuth
// @Param        task_id    query     int  false  "Filter by task ID"
// @Param        status     query     int  false  "Filter by status"
// @Param        page       query     int  false  "Page number (default: 1)"
// @Param        page_size  query     int  false  "Page size (default: 20, max: 100)"
// @Success      200  {object}  dto.TaskHistoryListResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /task-histories/my [get]
func (h *TaskHistoryHandler) GetMyHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "user not authenticated"})
		return
	}

	uid := userID.(int64)
	filter := &domain.TaskHistoryFilter{
		UserID:   &uid,
		Page:     1,
		PageSize: 20,
	}

	if taskID := c.Query("task_id"); taskID != "" {
		if id, err := strconv.ParseInt(taskID, 10, 64); err == nil {
			filter.TaskID = &id
		}
	}

	if status := c.Query("status"); status != "" {
		if s, err := strconv.Atoi(status); err == nil {
			filter.Status = &s
		}
	}

	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			filter.Page = p
		}
	}

	if pageSize := c.Query("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
			if ps > 100 {
				ps = 100
			}
			filter.PageSize = ps
		}
	}

	result, err := h.service.GetAllWithFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskHistoryListResponse(result, filter.Page, filter.PageSize))
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
