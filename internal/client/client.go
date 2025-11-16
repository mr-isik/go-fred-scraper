package client

import (
	"github.com/mr-isik/go-fred-scraper/pkg/config"
	"github.com/mr-isik/go-fred-scraper/pkg/fred"
)

// FredClient is an interface for the FRED client
type FredClient interface {
	GetSeriesObservations(seriesID, observationStart, observationEnd string) (*fred.ObservationsResponse, error)
}

// New returns a new FRED client
func New(cfg *config.Config) FredClient {
	return fred.NewClient(cfg.APIKey)
}
