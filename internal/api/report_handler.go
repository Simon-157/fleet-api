package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReportHandler handles report-related HTTP requests.
type ReportHandler struct {
	reportService service.ReportService
}

// NewReportHandler creates a new instance of ReportHandler.
func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

// GetDepartureAirportsWithInFlightAircraft handles the retrieval of departure airports with in-flight aircraft.
func (h *ReportHandler) GetDepartureAirportsWithInFlightAircraft(c *gin.Context) {
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	// Validate query parameters
	if startTime == "" || endTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_time and end_time query parameters are required"})
		return
	}

	// Call service method to get departure airports with in-flight aircraft
	departureAirports, err := h.reportService.GetFlightDetailsByTimeRange(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve departure airports with in-flight aircraft", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departureAirports)
}
