package service

import (
    "fleet_api/internal/model"
    "fleet_api/internal/repository"
    "github.com/pkg/errors"
    "time"
)

type FlightService struct {
    flightRepo repository.FlightRepository
}

func NewFlightService(flightRepo repository.FlightRepository) *FlightService {
    return &FlightService{
        flightRepo: flightRepo,
    }
}

func (s *FlightService) CreateFlight(flight *model.Flight) error {
	departureTime, err := time.Parse(time.RFC3339, flight.DepartureDatetime)
	if err != nil {
		return errors.Wrap(err, "failed to parse departure datetime")
	}

	if time.Now().After(departureTime) {
		return errors.New("departure datetime should be in the future")
	}

	err = s.flightRepo.CreateFlight(flight)
	if err != nil {
		return errors.Wrap(err, "failed to create flight")
	}
	return nil
}

func (s *FlightService) GetFlightByID(id int) (*model.Flight, error) {
    flight, err := s.flightRepo.GetFlightByID(id)
    if err != nil {
        return nil, errors.Wrap(err, "failed to get flight by ID")
    }
    return flight, nil
}

func (s *FlightService) UpdateFlight(flight *model.Flight) error {
    err := s.flightRepo.UpdateFlight(flight)
    if err != nil {
        return errors.Wrap(err, "failed to update flight")
    }
    return nil
}

// update flight aircraft id 

func (s *FlightService) UpdateFlightAircraft(flightID int, aircraftID int) error {
	err := s.flightRepo.UpdateFlightAircraft(flightID, aircraftID)
	if err != nil {
		return errors.Wrap(err, "failed to update flight aircraft")
	}
	return nil
}


func (s *FlightService) DeleteFlight(id int) error {
    err := s.flightRepo.DeleteFlight(id)
    if err != nil {
        return errors.Wrap(err, "failed to delete flight")
    }
    return nil
}


// search for flight given departure and arrival airports
func (s *FlightService) SearchFlightsByDepartureAndArrival(departureAirport string, arrivalAirport string) ([]model.FlightWithAircraft, error) {
    flights, err := s.flightRepo.SearchFlightsByAirport(departureAirport, arrivalAirport)
    if err != nil {
        return nil, errors.Wrap(err, "failed to search flights")
    }
    return flights, nil
}