package service

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
)

type FlightService struct {
	pv ports.FlightProviderContract
}

func NewFlightService(pv ports.FlightProviderContract) *FlightService {
	return &FlightService{
		pv: pv,
	}
}

func (svc *FlightService) RequestFlights(source, destination, departure string) ([]entities.Flight, error) {
	return svc.pv.RequestFlights(source, destination, departure)
}
