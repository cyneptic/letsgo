package provider

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
)

const (
	flightProviderHost     = "http://localhost:8000"
	flightProviderEndpoint = "/flights"
	httpTimeout            = 5 * time.Second
)

type SortFilterProviderClient struct {
	client  *http.Client
	address string
}

func NewSortFilterProviderClient(address string) *SortFilterProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{
		Transport: tr,
		Timeout:   httpTimeout,
	}

	return &SortFilterProviderClient{
		client: cl,
	}
}

func (pc *SortFilterProviderClient) RequestFlight(source, destination, departure string) ([]entities.Flight, error) {
	u, err := url.Parse(flightProviderHost + flightProviderEndpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("source", source)
	q.Set("destination", destination)
	q.Set("departing", departure)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var flights []entities.Flight
	err = json.NewDecoder(resp.Body).Decode(&flights)
	if err != nil {
		return nil, err
	}
	return flights, nil
}
