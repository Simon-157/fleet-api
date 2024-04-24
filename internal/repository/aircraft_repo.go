package repository

import (
	"database/sql"
	"fleet_api/internal/model"
	"github.com/pkg/errors"
)

type AircraftRepository struct {
    db *sql.DB
}

func NewAircraftRepository(db *sql.DB) *AircraftRepository {
    return &AircraftRepository{
        db: db,
    }
}

func (r *AircraftRepository) CreateAircraft(aircraft *model.Aircraft) error {
    _, err := r.db.Exec("INSERT INTO aircraft (serial_number, manufacturer) VALUES ($1, $2)",
        aircraft.SerialNumber, aircraft.Manufacturer)
    if err != nil {
        return errors.Wrap(err, "failed to create aircraft")
    }
    return nil
}

func (r *AircraftRepository) GetAircraftByID(id int) (*model.Aircraft, error) {
    var aircraft model.Aircraft
    err := r.db.QueryRow("SELECT id, serial_number, manufacturer FROM aircraft WHERE id = $1", id).
        Scan(&aircraft.ID, &aircraft.SerialNumber, &aircraft.Manufacturer)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.Wrap(err, "aircraft not found")
        }
        return nil, errors.Wrap(err, "failed to get aircraft")
    }
    return &aircraft, nil
}

func (r *AircraftRepository) UpdateAircraft(aircraft *model.Aircraft) error {
    _, err := r.db.Exec("UPDATE aircraft SET serial_number = $1, manufacturer = $2 WHERE id = $3",
        aircraft.SerialNumber, aircraft.Manufacturer, aircraft.ID)
    if err != nil {
        return errors.Wrap(err, "failed to update aircraft")
    }
    return nil
}

func (r *AircraftRepository) DeleteAircraft(id int) error {
    _, err := r.db.Exec("DELETE FROM aircraft WHERE id = $1", id)
    if err != nil {
        return errors.Wrap(err, "failed to delete aircraft")
    }
    return nil
}

// assign aircraft to flight
func (r *AircraftRepository) AssignAircraftToFlight(flightID int, aircraftID int) error {
	_, err := r.db.Exec("UPDATE flight SET aircraft_id = $1 WHERE id = $2", aircraftID, flightID)
	if err != nil {
		return errors.Wrap(err, "failed to assign aircraft to flight")
	}
	return nil
}


// check if flight exists
func (r *AircraftRepository) CheckFlightByID(flightID int) (*model.Flight, error) {
	var flight model.Flight
	err := r.db.QueryRow("SELECT id, departure_airport, arrival_airport, departure_datetime, arrival_datetime, aircraft_id FROM flight WHERE id = $1", flightID).
		Scan(&flight.ID, &flight.DepartureAirport, &flight.ArrivalAirport, &flight.DepartureDatetime, &flight.ArrivalDatetime, &flight.AircraftID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "flight not found")
		}
		return nil, errors.Wrap(err, "failed to get flight")
	}
	return &flight, nil
}