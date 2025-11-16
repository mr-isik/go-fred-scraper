package series

import (
	"github.com/mr-isik/go-fred-scraper/internal/client"
	"github.com/mr-isik/go-fred-scraper/pkg/fred"
)

// Service is a service for interacting with FRED series
type Service struct {
	client client.FredClient
}

// NewService creates a new series service
func NewService(client client.FredClient) *Service {
	return &Service{
		client: client,
	}
}

// GetSeriesObservations gets observations for a FRED series
func (s *Service) GetSeriesObservations(seriesID, observationStart, observationEnd string) (*fred.ObservationsResponse, error) {
	return s.client.GetSeriesObservations(seriesID, observationStart, observationEnd)
}
