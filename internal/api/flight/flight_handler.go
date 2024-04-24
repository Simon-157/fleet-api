package api

import (
	"fleet_api/internal/model"
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FlightHandler handles flight-related HTTP requests.
type FlightHandler struct {
	flightService service.FlightService
}

// NewFlightHandler creates a new instance of FlightHandler.
func NewFlightHandler(flightService service.FlightService) *FlightHandler {
	return &FlightHandler{
		flightService: flightService,
	}
}

// CreateFlight handles the creation of a new flight.
func (h *FlightHandler) CreateFlight(c *gin.Context) {
	var flight model.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.flightService.CreateFlight(&flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create flight"})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

// GetFlightByID handles the retrieval of a flight by ID.
func (h *FlightHandler) GetFlightByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	flight, err := h.flightService.GetFlightByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flight not found"})
		return
	}

	c.JSON(http.StatusOK, flight)
}

// UpdateFlight handles the update of an existing flight.
func (h *FlightHandler) UpdateFlight(c *gin.Context) {
	var flight model.Flight
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	flight.ID = id 

	if err := h.flightService.UpdateFlight(&flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update flight"})
		return
	}

	c.JSON(http.StatusOK, flight)
}

// DeleteFlight handles the deletion of an existing flight.
func (h *FlightHandler) DeleteFlight(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	if err := h.flightService.DeleteFlight(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete flight"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flight deleted successfully"})
}
