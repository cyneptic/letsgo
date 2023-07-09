package provider

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type TicketProviderClient struct {
	client  *http.Client
	address string
}

type TicketRequest struct {
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func NewTicketProviderClient() *TicketProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{Transport: tr}
	providerAddress := "http://localhost:8000/"

	return &TicketProviderClient{
		client:  cl,
		address: providerAddress,
	}
}

func (pc *TicketProviderClient) RequestCancelTicket(flightId uuid.UUID) error {
	// generate request payload
	payload, err := json.Marshal(ReserveRequest{
		Action: "cancel",
		Count:  1,
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
