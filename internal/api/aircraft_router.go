package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
)

type AircraftRouter struct {
	handler *AircraftHandler
}

func NewAircraftRouter(aircraftService service.AircraftService) *AircraftRouter {
	handler := NewAircraftHandler(aircraftService)
	return &AircraftRouter{handler: handler}
}

func (r *AircraftRouter) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/aircraft")
	group.POST("/", r.handler.CreateAircraft)
	group.GET("/:id", r.handler.GetAircraftByID)
	group.PUT("/:id", r.handler.UpdateAircraft)
	group.DELETE("/:id", r.handler.DeleteAircraft)

	// Error handler middleware
	router.Use(ErrorHandler())
}
