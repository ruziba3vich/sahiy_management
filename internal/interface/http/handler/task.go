package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appTask "github.com/ruziba3vich/sahiy_management/internal/application/task"
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

	task, err := h.service.Create(
		c.Request.Context(),
		req.ParentID,
		req.SectionID,
		req.Title,
		req.Description,
		req.Priority,
		req.Deadline,
		req.Status,
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
// @Summary      Get all tasks
// @Description  Retrieve a list of all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   dto.TaskResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /tasks [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTaskResponseList(tasks))
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
