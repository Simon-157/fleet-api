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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create aircraft"})
		return
	}

	c.JSON(http.StatusCreated, aircraft)
}

func (h *AircraftHandler) GetAircraftByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID"})
		return
	}

	aircraft, err := h.aircraftService.GetAircraftByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "aircraft not found"})
		return
	}

	c.JSON(http.StatusOK, aircraft)
}


func (h *AircraftHandler) UpdateAircraft(c *gin.Context) {
	var aircraft model.Aircraft
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID"})
		return
	}

	if err := c.ShouldBindJSON(&aircraft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	aircraft.ID = id 

	if err := h.aircraftService.UpdateAircraft(&aircraft); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update aircraft"})
		return
	}

	c.JSON(http.StatusOK, aircraft)
}


func (h *AircraftHandler) DeleteAircraft(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft ID"})
		return
	}

	if err := h.aircraftService.DeleteAircraft(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete aircraft"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "aircraft deleted successfully"})
}
