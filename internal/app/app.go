package app


import (
	"fleet_api/internal/api"
	"fleet_api/internal/repository"
	"fleet_api/internal/service"
	"fleet_api/config"
	"fleet_api/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize config
	cfg := config.LoadConfig()

	// Initialize storage
	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	// Initialize repositories
	aircraftRepo := repository.NewAircraftRepository(db.DB)
	flightRepo := repository.NewFlightRepository(db.DB)
	reportRepo := repository.NewReportRepository(db.DB)

	// Initialize services
	aircraftService := service.NewAircraftService(*aircraftRepo)
	flightService := service.NewFlightService(*flightRepo)
	reportService := service.NewReportService(*reportRepo)

	// Initialize routers
	aircraftRouter := api.NewAircraftRouter(*aircraftService)
	flightRouter := api.NewFlightRouter(*flightService)
	reportRouter := api.NewReportRouter(*reportService)

	// Register routes
	aircraftRouter.RegisterRoutes(router)
	flightRouter.RegisterRoutes(router)
	reportRouter.RegisterRoutes(router)

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Home route
	router.GET("/", func(c *gin.Context) {
	// API Endpoint Table
	endpointTable := `
	| Endpoint                   | Method | Description                        | Query Params                          | Required Body Content         |
	|----------------------------|--------|------------------------------------|---------------------------------------|-------------------------------|
	| /api/flight                | GET    | Retrieve all flights               | N/A                                   | N/A                           |
	| /api/flight/{id}           | GET    | Retrieve flight by ID              | N/A                                   | N/A                           |
	| /api/flight                | POST   | Create a new flight                | N/A                                   | Flight data                   |
	| /api/flight/{id}           | PUT    | Update flight details by ID        | N/A                                   | Updated flight data           |
	| /api/flight/{id}           | DELETE | Delete flight by ID                | N/A                                   | N/A                           |
	| /api/flight/search         | GET    | Search flights by departure and arrival airport | departure_airport, arrival_airport | N/A                           |
	| /api/aircraft              | GET    | Retrieve all aircrafts             | N/A                                   | N/A                           |
	| /api/aircraft/{id}         | GET    | Retrieve aircraft by ID            | N/A                                   | N/A                           |
	| /api/aircraft              | POST   | Create a new aircraft              | N/A                                   | Aircraft data                 |
	| /api/aircraft/{id}         | PUT    | Update aircraft details by ID     | N/A                                   | Updated aircraft data         |
	| /api/aircraft/{id}         | DELETE | Delete aircraft by ID              | N/A                                   | N/A                           |
	| /api/report/departure_airport | GET    | Retrieve all flights by departure airport | start_time, end_time                 | N/A                           |
	`
		c.String(200, "Welcome to the Fleet API!\n\n%s", endpointTable)
	})

	// Starting the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}