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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create flight", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

// GetFlights handles the retrieval of all flights.
func (h *FlightHandler) GetFlights(c *gin.Context) {
	flights, err := h.flightService.GetFlights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get flights", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flights)
}



// GetFlightByID handles the retrieval of a flight by ID.
func (h *FlightHandler) GetFlightByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID", "message": err.Error()})
		return
	}

	flight, err := h.flightService.GetFlightByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flight not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flight)
}

// UpdateFlight handles the update of an existing flight.
func (h *FlightHandler) UpdateFlight(c *gin.Context) {
	var flight model.Flight
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID", "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "message": err.Error()})
		return
	}

	flight.ID = id 

	if err := h.flightService.UpdateFlight(&flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update flight", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flight)
}

// UpdateFlightAircraftID handles the update of an existing flight's aircraft ID.
func (h *FlightHandler) UpdateFlightAircraftID(c *gin.Context) {
	// Parse aircraft ID from query parameter
	aircraftID, err := strconv.Atoi(c.Query("aircraft_id"))
	if err != nil || aircraftID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID"})
		return
	}

	// Parse flight ID from path parameter
	flightID, err := strconv.Atoi(c.Param("id"))
	if err != nil || flightID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	// Update flight's aircraft ID
	if err := h.flightService.UpdateFlightAircraft(flightID, aircraftID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update flight", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flight aircraft updated successfully"})
}


// DeleteFlight handles the deletion of an existing flight.
func (h *FlightHandler) DeleteFlight(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	// Validate query parameters
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flight ID"})
		return
	}

	if err := h.flightService.DeleteFlight(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete flight", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flight deleted successfully"})
}



// SearchFlightsByAirport handles the search of flights by departure and arrival airports.
func (h *FlightHandler) SearchFlightsByAirport(c *gin.Context) {
	departureAirport := c.Query("departure_airport")
	arrivalAirport := c.Query("arrival_airport")

	// Validate query parameters
	if departureAirport == "" || arrivalAirport == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "departure_airport and arrival_airport query parameters are required"})
		return
	}

	flights, err := h.flightService.SearchFlightsByDepartureAndArrival(departureAirport, arrivalAirport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search flights", "message": err.Error()})
		return
	}

	if len(flights) == 0 {
		c.JSON(http.StatusOK, gin.H{"flights": flights, "message": "no flights found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flights": flights, "message": "flights found"})
}
