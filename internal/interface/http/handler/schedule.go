package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appSchedule "github.com/ruziba3vich/sahiy_management/internal/application/schedule"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type ScheduleHandler struct {
	service *appSchedule.Service
}

func NewScheduleHandler(service *appSchedule.Service) *ScheduleHandler {
	return &ScheduleHandler{service: service}
}

// Create godoc
// @Summary      Create a new schedule
// @Description  Create a new schedule with the provided data
// @Tags         schedules
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateScheduleRequest  true  "Schedule data"
// @Success      201      {object}  dto.ScheduleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /schedules [post]
func (h *ScheduleHandler) Create(c *gin.Context) {
	var req dto.CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	status := 1
	if req.Status != nil {
		status = *req.Status
	}

	schedule, err := h.service.Create(c.Request.Context(), req.Name, req.Timezone, req.WeekDays, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToScheduleResponse(schedule))
}

// Update godoc
// @Summary      Update a schedule
// @Description  Update an existing schedule by ID
// @Tags         schedules
// @Accept       json
// @Produce      json
// @Param        id       path      int                        true  "Schedule ID"
// @Param        request  body      dto.UpdateScheduleRequest  true  "Schedule data"
// @Success      200      {object}  dto.ScheduleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /schedules/{id} [put]
func (h *ScheduleHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid schedule ID"})
		return
	}

	var req dto.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	schedule, err := h.service.Update(c.Request.Context(), id, req.Name, req.Timezone, req.WeekDays, req.Status)
	if err != nil {
		if errors.Is(err, postgres.ErrScheduleNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "schedule not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToScheduleResponse(schedule))
}

// Delete godoc
// @Summary      Delete a schedule
// @Description  Delete a schedule by ID
// @Tags         schedules
// @Produce      json
// @Param        id  path  int  true  "Schedule ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /schedules/{id} [delete]
func (h *ScheduleHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid schedule ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrScheduleNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "schedule not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a schedule by ID
// @Description  Retrieve a single schedule by its ID
// @Tags         schedules
// @Produce      json
// @Param        id  path      int  true  "Schedule ID"
// @Success      200  {object}  dto.ScheduleResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /schedules/{id} [get]
func (h *ScheduleHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid schedule ID"})
		return
	}

	schedule, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrScheduleNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "schedule not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToScheduleResponse(schedule))
}

// GetAll godoc
// @Summary      Get all schedules
// @Description  Retrieve a list of all schedules
// @Tags         schedules
// @Produce      json
// @Success      200  {array}   dto.ScheduleResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /schedules [get]
func (h *ScheduleHandler) GetAll(c *gin.Context) {
	schedules, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToScheduleResponseList(schedules))
}

func (h *ScheduleHandler) RegisterRoutes(r *gin.RouterGroup) {
	schedules := r.Group("/schedules")
	{
		schedules.POST("", h.Create)
		schedules.PUT("/:id", h.Update)
		schedules.DELETE("/:id", h.Delete)
		schedules.GET("/:id", h.GetByID)
		schedules.GET("", h.GetAll)
	}
}
