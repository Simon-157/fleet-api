package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FlightRouter handles flight-related routes.
type FlightRouter struct {
	handler *FlightHandler
}

// NewFlightRouter creates a new instance of FlightRouter.
func NewFlightRouter(flightService service.FlightService) *FlightRouter {
	handler := NewFlightHandler(flightService)
	return &FlightRouter{handler: handler}
}

// RegisterRoutes registers flight routes with the provided router.
func (r *FlightRouter) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/flight")
	group.POST("/", r.handler.CreateFlight)
	group.GET("/:id", r.handler.GetFlightByID)
	group.PUT("/:id", r.handler.UpdateFlight)
	group.DELETE("/:id", r.handler.DeleteFlight)

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
