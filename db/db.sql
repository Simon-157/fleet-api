CREATE TABLE aircraft (
    id SERIAL PRIMARY KEY,
    serial_number VARCHAR(50) NOT NULL,
    manufacturer VARCHAR(100) NOT NULL
);

CREATE TABLE flight (
    id SERIAL PRIMARY KEY,
    departure_airport VARCHAR(4) NOT NULL,
    arrival_airport VARCHAR(4) NOT NULL,
    departure_datetime TIMESTAMP NOT NULL,
    arrival_datetime TIMESTAMP NOT NULL,
    aircraft_id INT,
    FOREIGN KEY (aircraft_id) REFERENCES aircraft(id)
);
