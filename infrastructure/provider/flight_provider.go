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

type FlightProviderClient struct {
	client *http.Client
}

func NewFlightProviderClient() *FlightProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{
		Transport: tr,
		Timeout:   httpTimeout,
	}

	return &FlightProviderClient{
		client: cl,
	}
}

func (pc *FlightProviderClient) RequestFlights(source, destination, departure string) ([]entities.Flight, error) {
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

func (pc *FlightProviderClient) RequestFlight(id string) (entities.Flight, error) {
	var flight entities.Flight
	u, err := url.Parse(flightProviderHost + flightProviderEndpoint + "/" + id)
	if err != nil {
		return flight, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return flight, err

		resp, err := pc.client.Do(req)
		if err != nil {
			return flight, err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&flight)
		if err != nil {
			return flight, err
		}
		return flight, nil
	}
	resp, err := pc.client.Do(req)
	if err != nil {
		return flight, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&flight)
	if err != nil {
		return flight, err
	}
	return flight, nil
}
