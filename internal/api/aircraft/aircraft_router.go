package api

import (
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
