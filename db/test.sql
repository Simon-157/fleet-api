SELECT 
  departure_airport,
  COUNT(*) AS in_flight_aircraft,
  json_agg(json_build_object('aircraft_serial_number', aircraft.serial_number, 'in_flight_time' := flight_time)) AS in_flight_details
FROM (
  SELECT 
    f.departure_airport,
    a.serial_number,
    GREATEST(f.departure_datetime, '2024-05-01T00:00:00Z') AS departure_in_range,
    LEAST(f.arrival_datetime, '2024-05-31T23:59:59Z') AS arrival_in_range,
    EXTRACT(EPOCH FROM LEAST(f.arrival_datetime, '2024-05-31T23:59:59Z') - GREATEST(f.departure_datetime, '2024-05-01T00:00:00Z')) / 60 AS flight_time
  FROM flight f
  JOIN aircraft a ON f.aircraft_id = a.id
  JOIN (SELECT tst FROM UNNEST(TIMESTAMP tst WITH OFFSET '2024-05-01T00:00:00Z', ROWS PER RANGE INTERVAL 1 HOUR) AS dt(tst)) AS time_range
  ON (f.departure_datetime <= time_range.tst AND f.arrival_datetime >= time_range.tst)
  WHERE tst BETWEEN '2024-05-01T00:00:00Z' AND '2024-05-31T23:59:59Z'
) AS flight_data
GROUP BY departure_airport;
