package fred

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultAPIURL = "https://api.stlouisfed.org/fred"
)

// Client is a FRED API client
type Client struct {
	apiKey     string
	apiURL     string
	httpClient *http.Client
}

// NewClient creates a new FRED API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		apiURL:     defaultAPIURL,
		httpClient: &http.Client{},
	}
}

// GetSeriesObservations gets observations for a FRED series
func (c *Client) GetSeriesObservations(seriesID, observationStart, observationEnd string) (*ObservationsResponse, error) {
	endpoint := fmt.Sprintf("%s/series/observations", c.apiURL)
	params := url.Values{}
	params.Add("series_id", seriesID)
	params.Add("api_key", c.apiKey)
	params.Add("file_type", "json")
	params.Add("observation_start", observationStart)
	params.Add("observation_end", observationEnd)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var obsResp ObservationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&obsResp); err != nil {
		return nil, err
	}

	return &obsResp, nil
}

// Observation is a single data point for a FRED series
type Observation struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

// ObservationsResponse is the response from the FRED API for a series' observations
type ObservationsResponse struct {
	Observations []Observation `json:"observations"`
}
