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

	// Starting the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}