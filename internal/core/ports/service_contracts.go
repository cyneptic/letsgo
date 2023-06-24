package ports

import "github.com/cyneptic/letsgo/internal/core/entities"

type FlightServiceContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
}
