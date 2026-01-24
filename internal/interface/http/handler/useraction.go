package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	appUserAction "github.com/ruziba3vich/sahiy_management/internal/application/useraction"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type UserActionHandler struct {
	service *appUserAction.Service
}

func NewUserActionHandler(service *appUserAction.Service) *UserActionHandler {
	return &UserActionHandler{service: service}
}

// Create godoc
// @Summary      Create a new user action
// @Description  Create a new user action with the provided data
// @Tags         user-actions
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserActionRequest  true  "User action data"
// @Success      201      {object}  dto.UserActionResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /user-actions [post]
func (h *UserActionHandler) Create(c *gin.Context) {
	var req dto.CreateUserActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	action, err := h.service.Create(
		c.Request.Context(),
		req.UserID,
		req.VisitBranchID,
		req.LeaveBranchID,
		*req.ComeStatus,
		req.OutStatus,
		req.StartedAt,
		req.FinishedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToUserActionResponse(action))
}

// Update godoc
// @Summary      Update a user action
// @Description  Update an existing user action by ID
// @Tags         user-actions
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "User action ID"
// @Param        request  body      dto.UpdateUserActionRequest  true  "User action data"
// @Success      200      {object}  dto.UserActionResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /user-actions/{id} [put]
func (h *UserActionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid user action ID"})
		return
	}

	var req dto.UpdateUserActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	action, err := h.service.Update(
		c.Request.Context(),
		id,
		req.UserID,
		req.VisitBranchID,
		req.LeaveBranchID,
		*req.ComeStatus,
		req.OutStatus,
		req.StartedAt,
		req.FinishedAt,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrUserActionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "user action not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserActionResponse(action))
}

// Delete godoc
// @Summary      Delete a user action
// @Description  Delete a user action by ID
// @Tags         user-actions
// @Produce      json
// @Param        id  path  int  true  "User action ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /user-actions/{id} [delete]
func (h *UserActionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid user action ID"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrUserActionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "user action not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetByID godoc
// @Summary      Get a user action by ID
// @Description  Retrieve a single user action by its ID
// @Tags         user-actions
// @Produce      json
// @Param        id  path      int  true  "User action ID"
// @Success      200  {object}  dto.UserActionResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /user-actions/{id} [get]
func (h *UserActionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid user action ID"})
		return
	}

	action, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrUserActionNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "user action not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserActionResponse(action))
}

// GetAll godoc
// @Summary      Get all user actions
// @Description  Retrieve a list of all user actions
// @Tags         user-actions
// @Produce      json
// @Success      200  {array}   dto.UserActionResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /user-actions [get]
func (h *UserActionHandler) GetAll(c *gin.Context) {
	actions, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserActionResponseList(actions))
}

// Accept godoc
// @Summary      Accept a user action (check-in/check-out)
// @Description  Check-in or check-out based on user's geo-location relative to branches
// @Tags         user-actions
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AcceptUserActionRequest  true  "User location data"
// @Success      200      {object}  dto.UserActionResponse
// @Success      201      {object}  dto.UserActionResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /user-actions/accept [post]
func (h *UserActionHandler) Accept(c *gin.Context) {
	var req dto.AcceptUserActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	action, err := h.service.Accept(c.Request.Context(), req.UserID, req.Lat, req.Long)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserActionResponse(action))
}

// GetAttendance godoc
// @Summary      Get attendance reports
// @Description  Get attendance reports for a branch within a date range (Format: d.m.Y)
// @Tags         user-actions
// @Produce      json
// @Param        branch_id   query      int     false  "Branch ID (0 for all)"
// @Param        from_date   query      string  true   "From Date (Format: d.m.Y, e.g. 24.01.2026)"
// @Param        to_date     query      string  true   "To Date (Format: d.m.Y, e.g. 25.01.2026)"
// @Success      200         {array}    dto.UserAttendanceResponse
// @Failure      400         {object}   dto.ErrorResponse
// @Failure      500         {object}   dto.ErrorResponse
// @Router       /user-actions/attendance [get]
func (h *UserActionHandler) GetAttendance(c *gin.Context) {
	branchID, _ := strconv.ParseInt(c.Query("branch_id"), 10, 64)
	fromDateStr := c.Query("from_date")
	toDateStr := c.Query("to_date")

	const layout = "02.01.2006"
	fromTime, err := time.Parse(layout, fromDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid from_date format, expected d.m.Y"})
		return
	}
	toTime, err := time.Parse(layout, toDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid to_date format, expected d.m.Y"})
		return
	}

	// Set to end of day for toDate
	toTime = toTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	records, err := h.service.GetAttendance(c.Request.Context(), branchID, fromTime.Unix(), toTime.Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserAttendanceResponseList(records))
}

// GetStatuses godoc
// @Summary      Get user action status list
// @Description  Get a list of all possible user action statuses (ON_TIME, LATE, etc.)
// @Tags         user-actions
// @Produce      json
// @Success      200  {array}   dto.AttendanceStatusResponse
// @Router       /user-actions/statuses [get]
func (h *UserActionHandler) GetStatuses(c *gin.Context) {
	labels := h.service.GetStatuses()
	var resp []dto.AttendanceStatusResponse
	for id, label := range labels {
		resp = append(resp, dto.AttendanceStatusResponse{
			ID:    id,
			Label: label,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserActionHandler) RegisterRoutes(r *gin.RouterGroup) {
	actions := r.Group("/user-actions")
	{
		actions.POST("", h.Create)
		actions.POST("/accept", h.Accept)
		actions.GET("/attendance", h.GetAttendance)
		actions.GET("/statuses", h.GetStatuses)
		actions.PUT("/:id", h.Update)
		actions.DELETE("/:id", h.Delete)
		actions.GET("/:id", h.GetByID)
		actions.GET("", h.GetAll)
	}
}
