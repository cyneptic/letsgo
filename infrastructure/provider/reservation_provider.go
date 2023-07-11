package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

var (
	ErrReservationFailed = errors.New("Failed to make reservation, Error from provider")
)

type ReservationProviderClient struct {
	client  *http.Client
	address string
}

type ReserveRequest struct {
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func NewReservationProviderClient() *ReservationProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{Transport: tr}

	providerAddress := "http://localhost:8000/"
	return &ReservationProviderClient{
		client:  cl,
		address: providerAddress,
	}
}

func (pc *ReservationProviderClient) RequestFlightByID(flightId uuid.UUID) (entities.Flight, error) {
	resp, err := pc.client.Get(pc.address + "/flights/?id=" + flightId.String())
	if err != nil {
		return entities.Flight{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entities.Flight{}, err
	}

	var response entities.Flight
	if err = json.Unmarshal(body, &response); err != nil {
		return entities.Flight{}, err
	}

	return response, nil
}

func (pc *ReservationProviderClient) RequestReserve(flightId uuid.UUID, count int) error {
	// generate request payload
	payload, err := json.Marshal(ReserveRequest{
		Action: "reserve",
		Count:  count,
	})
	if err != nil {
		return err
	}

	// setup url for request
	url, err := url.Parse(pc.address)
	if err != nil {
		return err
	}
	q := url.Query()
	q.Set("id", flightId.String())
	url.RawQuery = q.Encode()

	// make a patch request using the payload and the address of provider
	req, err := http.NewRequest("PATCH", url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil
	}

	//send the request with content-type application/json header
	req.Header.Set("Content-Type", "application/json")
	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read the response and return an error if there is a problem or return nil if reservation is made
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response bool
	if err = json.Unmarshal(body, &response); err != nil {
		return err
	}
	if response == false {
		return ErrReservationFailed
	}
	return nil
}

func (pc *ReservationProviderClient) RequestCancelReservation(flightId uuid.UUID, count int) error {
	// generate request payload
	payload, err := json.Marshal(ReserveRequest{
		Action: "cancel",
		Count:  count,
	})
	if err != nil {
		return err
	}

	// setup url for request
	url, err := url.Parse(pc.address)
	if err != nil {
		return err
	}
	q := url.Query()
	q.Set("id", flightId.String())
	url.RawQuery = q.Encode()

	// make a patch request using the payload and the address of provider
	req, err := http.NewRequest("PATCH", url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil
	}

	//send the request with content-type application/json header
	req.Header.Set("Content-Type", "application/json")
	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read the response and return an error if there is a problem or return nil if reservation has been cancelled
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response bool
	if err = json.Unmarshal(body, &response); err != nil {
		return err
	}
	if response == false {
		return ErrReservationFailed
	}
	return nil
}
