package ports

import (
	"github.com/cyneptic/letsgo/internal/domain/entities"
	"github.com/google/uuid"
)

type ReserveServiceContract interface {
	CancelReservation(rId uuid.UUID) error
	GetUserReservations(userId uuid.UUID) ([]entities.Reservation, error)
	GetAllReservations() ([]entities.Reservation, error)
	Reserve(flightId uuid.UUID, userId uuid.UUID, passengers []uuid.UUID) (uuid.UUID, error)
}

type TicketServiceContract interface {
	IsCancellable(ticketId uuid.UUID) (bool, error)
	CancelTicket(ticketId uuid.UUID) error
}
