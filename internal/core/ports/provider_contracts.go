package ports

import (
	"github.com/cyneptic/letsgo/internal/domain/entities"
	"github.com/google/uuid"
)

type ReserveProviderContract interface {
	RequestFlightByID(flightId uuid.UUID) (entities.Flight, error)
	RequestReserve(flightId uuid.UUID, count int) error
	RequestCancelReservation(flightId uuid.UUID, count int) error
}

type TicketProviderContract interface {
	RequestCancelTicket(flightId uuid.UUID) error
}
