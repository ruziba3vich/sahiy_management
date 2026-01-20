package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appTask "github.com/ruziba3vich/sahiy_management/internal/application/task"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type TaskHandler struct {
	service *appTask.Service
}

func NewTaskHandler(service *appTask.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

// Create godoc
// @Summary      Create a new task
// @Description  Create a new task with the provided data
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.CreateTaskRequest  true  "Task data"
// @Success      201      {object}  dto.TaskResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	status := 1
	if req.Status != nil {
		status = *req.Status
	}

	task, err := h.service.Create(
		c.Request.Context(),
		req.ParentID,
		req.SectionID,
		req.UserID,
		req.Title,
		req.Description,
		req.Priority,
		req.Deadline,
		status,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToTaskResponse(task))
}

// Update godoc
// @Summary      Update a task
// @Description  Update an existing task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int                    true  "Task ID"
// @Param        request  body      dto.UpdateTaskRequest  true  "Task data"
// @Success      200      {object}  dto.TaskResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /tasks/{id} [put]
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task ID"})
		return
	}

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	task, err := h.service.Update(
		c.Request.Context(),
		id,
		req.ParentID,
		req.SectionID,
		req.UserID,
		req.Title,
		req.Description,
		req.Priority,
		req.Deadline,
		req.Status,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskResponse(task))
}

// Delete godoc
// @Summary      Delete a task
// @Description  Delete a task by ID
// @Tags         tasks
// @Produce      json
// @Param        id  path  int  true  "Task ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /tasks/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a task by ID
// @Description  Retrieve a single task by its ID
// @Tags         tasks
// @Produce      json
// @Param        id  path      int  true  "Task ID"
// @Success      200  {object}  dto.TaskResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /tasks/{id} [get]
func (h *TaskHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task ID"})
		return
	}

	task, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskResponse(task))
}

// GetAll godoc
// @Summary      Get all tasks with filtering and pagination
// @Description  Retrieve a list of tasks with optional filters (user_id, section_id, status, priority, parent_id) and pagination
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        user_id     query     int  false  "Filter by user ID"
// @Param        section_id  query     int  false  "Filter by section ID"
// @Param        status      query     int  false  "Filter by status"
// @Param        priority    query     int  false  "Filter by priority"
// @Param        parent_id   query     int  false  "Filter by parent task ID"
// @Param        page        query     int  false  "Page number (default: 1)"
// @Param        page_size   query     int  false  "Page size (default: 20, max: 100)"
// @Success      200  {object}  dto.TaskListResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /tasks [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	filter := &domain.TaskFilter{
		Page:     1,
		PageSize: 20,
	}

	if userID := c.Query("user_id"); userID != "" {
		if id, err := strconv.ParseInt(userID, 10, 64); err == nil {
			filter.UserID = &id
		}
	}

	if sectionID := c.Query("section_id"); sectionID != "" {
		if id, err := strconv.ParseInt(sectionID, 10, 64); err == nil {
			filter.SectionID = &id
		}
	}

	if status := c.Query("status"); status != "" {
		if s, err := strconv.Atoi(status); err == nil {
			filter.Status = &s
		}
	}

	if priority := c.Query("priority"); priority != "" {
		if p, err := strconv.Atoi(priority); err == nil {
			filter.Priority = &p
		}
	}

	if parentID := c.Query("parent_id"); parentID != "" {
		if id, err := strconv.ParseInt(parentID, 10, 64); err == nil {
			filter.ParentID = &id
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

	c.JSON(http.StatusOK, dto.ToTaskListResponse(result, filter.Page, filter.PageSize))
}

// GetMyTasks godoc
// @Summary      Get current user's tasks with filtering and pagination
// @Description  Retrieve tasks assigned to the authenticated user with optional filters (section_id, status, priority, parent_id) and pagination
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        section_id  query     int  false  "Filter by section ID"
// @Param        status      query     int  false  "Filter by status"
// @Param        priority    query     int  false  "Filter by priority"
// @Param        parent_id   query     int  false  "Filter by parent task ID"
// @Param        page        query     int  false  "Page number (default: 1)"
// @Param        page_size   query     int  false  "Page size (default: 20, max: 100)"
// @Success      200  {object}  dto.TaskListResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /tasks/my [get]
func (h *TaskHandler) GetMyTasks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "user not authenticated"})
		return
	}

	uid := userID.(int64)
	filter := &domain.TaskFilter{
		UserID:   &uid,
		Page:     1,
		PageSize: 20,
	}

	if sectionID := c.Query("section_id"); sectionID != "" {
		if id, err := strconv.ParseInt(sectionID, 10, 64); err == nil {
			filter.SectionID = &id
		}
	}

	if status := c.Query("status"); status != "" {
		if s, err := strconv.Atoi(status); err == nil {
			filter.Status = &s
		}
	}

	if priority := c.Query("priority"); priority != "" {
		if p, err := strconv.Atoi(priority); err == nil {
			filter.Priority = &p
		}
	}

	if parentID := c.Query("parent_id"); parentID != "" {
		if id, err := strconv.ParseInt(parentID, 10, 64); err == nil {
			filter.ParentID = &id
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

	c.JSON(http.StatusOK, dto.ToTaskListResponse(result, filter.Page, filter.PageSize))
}

func (h *TaskHandler) RegisterRoutes(r *gin.RouterGroup) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("", h.Create)
		tasks.PUT("/:id", h.Update)
		tasks.DELETE("/:id", h.Delete)
		tasks.GET("/:id", h.GetByID)
		tasks.GET("", h.GetAll)
	}
}
