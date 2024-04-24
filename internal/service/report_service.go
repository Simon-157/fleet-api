package service

import (
    "fleet_api/internal/model"
    "fleet_api/internal/repository"
    "github.com/pkg/errors"
)

type ReportService struct {
    reportRepo repository.ReportRepository
}

func NewReportService(reportRepo repository.ReportRepository) *ReportService {
    return &ReportService{
        reportRepo: reportRepo,
    }
}

// get flight details for report given start and end time

func (s *ReportService) GetFlightDetailsByTimeRange(startDatetime, endDatetime string) ([]model.FlightDetail, error) {
    flightDetails, err := s.reportRepo.GetFlightDetailsByTimeRange(startDatetime, endDatetime)
    if err != nil {
        return nil, errors.Wrap(err, "failed to get flight report details")
    }
    return flightDetails, nil
}