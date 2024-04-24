package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReportRouter handles report-related routes.
type ReportRouter struct {
	handler *ReportHandler
}

// NewReportRouter creates a new instance of ReportRouter.
func NewReportRouter(reportService service.ReportService) *ReportRouter {
	handler := NewReportHandler(reportService)
	return &ReportRouter{handler: handler}
}

// RegisterRoutes registers report routes with the provided router.
func (r *ReportRouter) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/report")
	group.GET("/departure_airports", r.handler.GetDepartureAirportsWithInFlightAircraft)

	// Error handler middleware
	router.Use(ErrorHandler())
}

// ErrorHandler is a middleware for handling errors and returning appropriate responses.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}
	}
}
