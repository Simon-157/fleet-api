package api

import (
	"fleet_api/internal/model"
	"fleet_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AircraftHandler struct {
	aircraftService service.AircraftService
}

func NewAircraftHandler(aircraftService service.AircraftService) *AircraftHandler {
	return &AircraftHandler{
		aircraftService: aircraftService,
	}
}

func (h *AircraftHandler) CreateAircraft(c *gin.Context) {
	var aircraft model.Aircraft

	// GET FLIGHT ID FROM URL IF ANY ELSE DEFAULT TO 0
	flightID, _ := strconv.Atoi(c.DefaultQuery("flight_id", "0"))
	
	if err := c.ShouldBindJSON(&aircraft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.aircraftService.CreateAircraft(&aircraft, flightID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create aircraft", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, aircraft)
}

//get all aircrafts
func (h *AircraftHandler) GetAircrafts(c *gin.Context) {
	aircrafts, err := h.aircraftService.GetAircrafts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get aircrafts", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aircrafts)
}

// GetAircraftByID retrieves an aircraft by its ID
func (h *AircraftHandler) GetAircraftByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID", "message": err.Error()})
		return
	}

	aircraft, err := h.aircraftService.GetAircraftByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "aircraft not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aircraft)
}


// UpdateAircraft updates an existing aircraft
func (h *AircraftHandler) UpdateAircraft(c *gin.Context) {
	var aircraft model.Aircraft
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID", "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&aircraft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "message": err.Error()})
		return
	}

	aircraft.ID = id 

	if err := h.aircraftService.UpdateAircraft(&aircraft); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update aircraft", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aircraft)
}


// DeleteAircraft deletes an existing aircraft
func (h *AircraftHandler) DeleteAircraft(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID"})
		return
	}

	if err := h.aircraftService.DeleteAircraft(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete aircraft", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "aircraft deleted successfully"})
}
