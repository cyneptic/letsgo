package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type ReserveProviderContract interface {
	RequestFlightByID(flightId uuid.UUID) (entities.Flight, error)
	RequestReserve(flightId uuid.UUID, count int) error
	RequestCancelReservation(flightId uuid.UUID, count int) error
}

type SortFilterProviderContract interface {
	RequestFlight(source, destination, departure string) ([]entities.Flight, error)
}
