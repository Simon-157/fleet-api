package service

import (
	"fleet_api/internal/repository"
)

type Service interface {
}

func NewService(aircraftRepo repository.AircraftRepository, flightRepo repository.FlightRepository, reportRepo repository.ReportRepository) Service {
    return &service{
        Aircraft: NewAircraftService(aircraftRepo),
        Flight:   NewFlightService(flightRepo),
        Report:   NewReportService(reportRepo),
    }
}

type service struct {
    Aircraft *AircraftService
    Flight   *FlightService
    Report   *ReportService
}
