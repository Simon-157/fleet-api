package service

import (
    "fleet_api/internal/model"
    "fleet_api/internal/repository"
    "github.com/pkg/errors"
)

type AircraftService struct {
    aircraftRepo repository.AircraftRepository
}

func NewAircraftService(aircraftRepo repository.AircraftRepository) *AircraftService {
    return &AircraftService{
        aircraftRepo: aircraftRepo,
    }
}

func (s *AircraftService) CreateAircraft(aircraft *model.Aircraft, flightId int) error {
    err := s.aircraftRepo.CreateAircraft(aircraft)
    if err != nil {
        return errors.Wrap(err, "failed to create aircraft")
    }

	//after creating aircraft, check if flightid was provided
	if flightId != 0 {

		// check if flight exists
		_, err := s.aircraftRepo.CheckFlightByID(flightId)
		if err != nil {
			return errors.Wrap(err, "flight not found")
		}
		// assign aircraft to flight
		err = s.aircraftRepo.AssignAircraftToFlight(flightId, aircraft.ID)
		if err != nil {
			return errors.Wrap(err, "failed to assign aircraft to flight")
		}
	
	}
    return nil
}

func (s *AircraftService) GetAircraftByID(id int) (*model.Aircraft, error) {
    aircraft, err := s.aircraftRepo.GetAircraftByID(id)
    if err != nil {
        return nil, errors.Wrap(err, "failed to get aircraft by ID")
    }
    return aircraft, nil
}

func (s *AircraftService) UpdateAircraft(aircraft *model.Aircraft) error {
    err := s.aircraftRepo.UpdateAircraft(aircraft)
    if err != nil {
        return errors.Wrap(err, "failed to update aircraft")
    }
    return nil
}

func (s *AircraftService) DeleteAircraft(id int) error {
    err := s.aircraftRepo.DeleteAircraft(id)
    if err != nil {
        return errors.Wrap(err, "failed to delete aircraft")
    }
    return nil
}

// func (s *AircraftService) GetAircrafts() ([]model.Aircraft, error) {
// 	aircrafts, err := s.aircraftRepo.GetAircrafts()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get aircrafts")
// 	}
// 	return aircrafts, nil
// }