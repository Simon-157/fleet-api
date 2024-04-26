
# Fleet API

Welcome to the Fleet API project! This API serves as the backend for managing fleet-related operations.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Docker](#docker)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This API provides functionalities for managing flights, aircraft, and reports within a fleet management system.

## Installation

To install and run the Fleet API locally, follow these steps:

1. Clone this repository:
   ```bash
   git clone https://github.com/Simon-157/fleet-api.git
   ```

2. Navigate to the project directory:
   ```bash
   cd fleet_api
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Build the application:
   ```bash
   go build -o fleet-api cmd/fleet-api/main.go
   ```

5. Run the application:
   ```bash
   ./fleet-api
   ```

## Usage

Once the application is running, you can access the API endpoints using a tool like Postman or by making HTTP requests from your application.

For detailed API documentation, refer to the [API Endpoints](#api-endpoints) section below.

## Docker

Alternatively, you can run the Fleet API using Docker. Follow these steps:

1. Build the Docker image:
   ```bash
   docker build -t fleet-api .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 5000:5000 fleet-api
   ```

## API Endpoints
## API Endpoints

| Endpoint                   | Method | Description                         | Query Params                          | Required Body Content                 |
|----------------------------|--------|-------------------------------------|---------------------------------------|--------------------------------------|
| `/api/flight`             | GET    | Retrieve all flights                | N/A                                   | N/A                                  |
| `/api/flight/{id}`        | GET    | Retrieve flight by ID               | N/A                                   | N/A                                  |
| `/api/flight`             | POST   | Create a new flight   (provide the an id as search params when creating with an assigned aircraft)             | N/A                                   | Flight data                         |
| `/api/flight/{id}`        | PUT    | Update flight details by ID         | N/A                                   | Updated flight data                 |
| `/api/flight/{id}`        | DELETE | Delete flight by ID                 | N/A                                   | N/A                                  |
| `/api/flight/search`            | GET    | search by departure_airport,arrival_airport              | departure_airport,arrival_airport                                    | N/A                                  |
| `/api/aircraft`           | GET    | Retrieve all aircrafts              | N/A                                   | N/A                                  |
| `/api/aircraft/{id}`      | GET    | Retrieve aircraft by ID             | N/A                                   | N/A                                  |
| `/api/aircraft`           | POST   | Create a new aircraft               | N/A                                   | Aircraft data                       |
| `/api/aircraft/{id}`      | PUT    | Update aircraft details by ID      | N/A                                   | Updated aircraft data               |
| `/api/aircraft/{id}`      | DELETE | Delete aircraft by ID               | N/A                                   | N/A                                  |
| `/api/report/departure_airport`             | GET    | Retrieve all                |start_time and end_time                                 | N/A                                  |

## Sample Testing with Postman
## API Endpoints

### Aircraft Endpoints:

#### Create Aircraft:

- **Method:** POST
- **URL:** http://localhost:8080/api/aircraft
- **Body:** JSON
  ```json
  {
      "serial_number": "JKL013",
      "manufacturer": "Embraer"
  }
  ```

#### Create Flight with Aircraft ID:

- **Method:** POST
- **URL:** http://localhost:8080/api/flight
- **Body:** JSON
  ```json
  {
      "departure_airport": "LHR",
      "arrival_airport": "ICN",
      "departure_datetime": "2024-05-24T08:00:00Z",
      "arrival_datetime": "2024-05-24T15:00:00Z",
      "aircraft_id": 2
  }
  ```

#### Create Flight without Aircraft ID:

- **Method:** POST
- **URL:** http://localhost:8080/api/flight
- **Body:** JSON
  ```json
  {
      "departure_airport": "IST",
      "arrival_airport": "NRT",
      "departure_datetime": "2024-05-26T08:00:00Z",
      "arrival_datetime": "2024-05-26T18:00:00Z"
  }
  ```

#### Search Flights Route:

- **Method:** GET
- **URL:** http://localhost:8080/api/flight/search
- **Params:**
  - `departure_airport`: Departure airport code (e.g., LHR)
  - `arrival_airport`: Arrival airport code (e.g., JFK)
- **Example:**
  Searching for flights from London Heathrow Airport (LHR) to John F. Kennedy International Airport (JFK):
  URL: http://localhost:8080/api/flight/search?departure_airport=LHR&arrival_airport=JFK

#### Reports Route:

- **Method:** GET
- **URL:** http://localhost:8080/api/report/departure_airports
- **Params:**
  - `start_time`: Start time for the search range (e.g., 2024-05-01T00:00:00Z)
  - `end_time`: End time for the search range (e.g., 2024-05-31T23:59:59Z)
- **Example:**
  Get departure airports with in-flight aircraft for the month of May 2024:
  URL: http://localhost:8080/api/report/departure_airports?start_time=2024-05-01T00:00:00Z&end_time=2024-05-31T23:59:59Z


## Contributing

Contributions to the Fleet API project are welcome! To contribute, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature/your-feature`)
6. Create a new pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```