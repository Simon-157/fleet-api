package model

import "encoding/json"

type Aircraft struct {
    ID            int    `json:"id"`
    SerialNumber  string `json:"serial_number"`
    Manufacturer  string `json:"manufacturer"`
}

type Flight struct {
    ID               int    `json:"id"`
    DepartureAirport string `json:"departure_airport"`
    ArrivalAirport   string `json:"arrival_airport"`
    DepartureDatetime string `json:"departure_datetime"`
    ArrivalDatetime string `json:"arrival_datetime"`
    AircraftID       int    `json:"aircraft_id"`
}

type FlightWithAircraft struct {
	Flight *Flight
    Aircraft *Aircraft
}

type Report struct {
    DepartureAirport string            `json:"departure_airport"`
    InFlightAircraft int               `json:"in_flight_aircraft"`
    AircraftList     []AircraftInFlight `json:"aircraft_list"`
}



type AircraftInFlight struct {
    AircraftSerialNumber string `json:"aircraft_serial_number"`
    InFlightTime         json.RawMessage `json:"in_flight_time"`
}



type FlightDetail struct {
	DepartureAirport string `json:"departure_airport"`
	InFlightAircraft int    `json:"in_flight_aircraft"`
	AircraftDetails  string `json:"aircraft_details"`
}