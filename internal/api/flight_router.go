package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"

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

func (r *FlightRouter) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/flight")
	group.POST("/", r.handler.CreateFlight)
	group.GET("/", r.handler.GetFlights)
	group.GET("/:id", r.handler.GetFlightByID)
	group.PUT("/:id", r.handler.UpdateFlight)
	group.DELETE("/:id", r.handler.DeleteFlight)
	group.GET("/search", r.handler.SearchFlightsByAirport)
	// update flight aircraft
	group.PUT("/:id/aircraft", r.handler.UpdateFlightAircraftID)

	// Error handler middleware
	router.Use(ErrorHandler())
}
