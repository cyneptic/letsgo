package provider

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/cyneptic/letsgo/internal/core/entities"
)

type SortFilterProviderClient struct {
	client  *http.Client
	address string
}

func NewSortFilterProviderClient(address string) *SortFilterProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{Transport: tr}

	return &SortFilterProviderClient{
		client:  cl,
		address: address,
	}
}

func (pc *SortFilterProviderClient) RequestFlight() ([]entities.Flight, error) {
	resp, err := pc.client.Get(pc.address + "/flights?source=tehran&destination=sari&departing=2023-07-29")
	if err != nil {
		return []entities.Flight{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []entities.Flight{}, err
	}
	var response []entities.Flight
	if err = json.Unmarshal(body, &response); err != nil {
		return []entities.Flight{}, err
	}

	return response, nil
}
