package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appStatistics "github.com/ruziba3vich/sahiy_management/internal/application/statistics"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/dto"
)

type StatisticsHandler struct {
	service *appStatistics.Service
}

func NewStatisticsHandler(service *appStatistics.Service) *StatisticsHandler {
	return &StatisticsHandler{service: service}
}

// GetDashboardStats godoc
// @Summary      Get dashboard statistics
// @Description  Retrieve total employees, department stats, total attendance, and task stats for the current month
// @Tags         statistics
// @Produce      json
// @Success      200  {object}  statistics.DashboardStats
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /statistics/dashboard [get]
func (h *StatisticsHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.service.GetDashboardStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *StatisticsHandler) RegisterRoutes(r *gin.RouterGroup) {
	stats := r.Group("/statistics")
	{
		stats.GET("/dashboard", h.GetDashboardStats)
	}
}
