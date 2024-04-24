package repository

import (
	"database/sql"
	"fleet_api/internal/model"

	"github.com/pkg/errors"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{
		db: db,
	}
}


func (r *ReportRepository) GetFlightDetailsByTimeRange(startDatetime, endDatetime string) ([]model.FlightDetail, error) {
	query := `
		SELECT 
			departure_airport,
			COUNT(*) FILTER (WHERE departure_datetime BETWEEN $1 AND $2) AS in_flight_aircraft,
			json_agg(
				json_build_object(
					'aircraft_serial_number', aircraft.serial_number,
					'in_flight_time',
					CASE 
						WHEN departure_datetime < $1 THEN 
							(interval '1 minute' * (extract(epoch FROM $2 - departure_datetime)))
						WHEN arrival_datetime > $2 THEN 
							(interval '1 minute' * (extract(epoch FROM arrival_datetime - $2)))
						ELSE
							interval '1 minute' * (extract(epoch FROM arrival_datetime - departure_datetime))
					END
				)
			) AS aircraft_details
		FROM flight
		INNER JOIN aircraft ON flight.aircraft_id = aircraft.id
		WHERE departure_datetime BETWEEN $1 AND $2 
		OR (departure_datetime < $1 AND arrival_datetime > $2)
		GROUP BY departure_airport;
	`
	rows, err := r.db.Query(query, startDatetime, endDatetime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}
	defer rows.Close()

	var flightDetails []model.FlightDetail
	for rows.Next() {
		var detail model.FlightDetail
		if err := rows.Scan(&detail.DepartureAirport, &detail.InFlightAircraft, &detail.AircraftDetails); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		flightDetails = append(flightDetails, detail)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error in rows")
	}
	return flightDetails, nil
}
