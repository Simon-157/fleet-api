package repository

import (
    "database/sql"
    "fleet_api/internal/model"
    "github.com/pkg/errors"
)

type FlightRepository struct {
    db *sql.DB
}

func NewFlightRepository(db *sql.DB) *FlightRepository {
    return &FlightRepository{
        db: db,
    }
}

func (r *FlightRepository) CreateFlight(flight *model.Flight) error {
	// TODO: add validation

	switch {
		case flight.AircraftID == 0:
			_, err := r.db.Exec("INSERT INTO flight (departure_airport, arrival_airport, departure_datetime, arrival_datetime) VALUES ($1, $2, $3, $4)",
				flight.DepartureAirport, flight.ArrivalAirport, flight.DepartureDatetime, flight.ArrivalDatetime)
			if err != nil {
				return errors.Wrap(err, "failed to create flight")
			}
			return nil

		case flight.AircraftID != 0:
			_, err := r.db.Exec("INSERT INTO flight (departure_airport, arrival_airport, departure_datetime, arrival_datetime, aircraft_id) VALUES ($1, $2, $3, $4, $5)",
				flight.DepartureAirport, flight.ArrivalAirport, flight.DepartureDatetime, flight.ArrivalDatetime, flight.AircraftID)
			if err != nil {
				return errors.Wrap(err, "failed to create flight")
			}
			return nil
		default:
			return errors.New("invalid flight object")
	}
	
}

func (r *FlightRepository) GetFlights() ([]model.Flight, error) {
rows, err := r.db.Query("SELECT id, departure_airport, arrival_airport, departure_datetime, arrival_datetime, aircraft_id FROM flight")
if err != nil {
	return nil, errors.Wrap(err, "failed to get flights")
}
defer rows.Close()
var flights []model.Flight
for rows.Next() {
	var flight model.Flight
	var aircraftID sql.NullInt64 // Use sql.NullInt64 to handle NULL values
	if err := rows.Scan(&flight.ID, &flight.DepartureAirport, &flight.ArrivalAirport, &flight.DepartureDatetime, &flight.ArrivalDatetime, &aircraftID); err != nil {
		return nil, errors.Wrap(err, "failed to scan flight")
	}
	flight.AircraftID = int(aircraftID.Int64) // Convert int64 to int
	flights = append(flights, flight)
}
    return flights, nil
}


func (r *FlightRepository) GetFlightByID(id int) (*model.Flight, error) {
    var flight model.Flight
    err := r.db.QueryRow("SELECT id, departure_airport, arrival_airport, departure_datetime, arrival_datetime, aircraft_id FROM flight WHERE id = $1", id).
        Scan(&flight.ID, &flight.DepartureAirport, &flight.ArrivalAirport, &flight.DepartureDatetime, &flight.ArrivalDatetime, &flight.AircraftID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.Wrap(err, "flight not found")
        }
        return nil, errors.Wrap(err, "failed to get flight")
    }
    return &flight, nil
}

func (r *FlightRepository) UpdateFlight(flight *model.Flight) error {
    _, err := r.db.Exec("UPDATE flight SET departure_airport = $1, arrival_airport = $2, departure_datetime = $3, arrival_datetime = $4, aircraft_id = $5 WHERE id = $6",
        flight.DepartureAirport, flight.ArrivalAirport, flight.DepartureDatetime, flight.ArrivalDatetime, flight.AircraftID, flight.ID)
    if err != nil {
        return errors.Wrap(err, "failed to update flight")
    }
    return nil
}

// update flight aircraft
func (r *FlightRepository) UpdateFlightAircraft(flightID int, aircraftID int) error {
	_, err := r.db.Exec("UPDATE flight SET aircraft_id = $1 WHERE id = $2", aircraftID, flightID)
	if err != nil {
		return errors.Wrap(err, "failed to update flight aircraft")
	}
	return nil
}

func (r *FlightRepository) DeleteFlight(id int) error {
    _, err := r.db.Exec("DELETE FROM flight WHERE id = $1", id)
    if err != nil {
        return errors.Wrap(err, "failed to delete flight")
    }
    return nil
}


func (r *FlightRepository) SearchFlightsByAirport(departureAirport string, arrivalAirport string) ([]model.FlightWithAircraft, error) {
	rows, err := r.db.Query(`
        SELECT f.id, f.departure_airport, f.arrival_airport, f.departure_datetime, f.arrival_datetime, f.aircraft_id,
               a.serial_number, a.manufacturer
        FROM flight f
        LEFT JOIN aircraft a ON f.aircraft_id = a.id
        WHERE f.departure_airport = $1 AND f.arrival_airport = $2
    `, departureAirport, arrivalAirport)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch flights by airports")
	}
	
	defer rows.Close()
	var flights []model.FlightWithAircraft
	
	for rows.Next() {
		var flightWithAircraft model.FlightWithAircraft
		flightWithAircraft.Flight = &model.Flight{}

		var aircraftID sql.NullInt64
		var serialNumber, manufacturer sql.NullString
		if err := rows.Scan(
			&flightWithAircraft.Flight.ID,
			&flightWithAircraft.Flight.DepartureAirport,
			&flightWithAircraft.Flight.ArrivalAirport,
			&flightWithAircraft.Flight.DepartureDatetime,
			&flightWithAircraft.Flight.ArrivalDatetime,
			&aircraftID,
			&serialNumber,
			&manufacturer,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		if aircraftID.Valid {
			flightWithAircraft.Aircraft = &model.Aircraft{
				ID:           int(aircraftID.Int64),
				SerialNumber: serialNumber.String,
				Manufacturer: manufacturer.String,
			}
		}
		flights = append(flights, flightWithAircraft)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error in rows")
	}
	return flights, nil
}