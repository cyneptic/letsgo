package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type ReserveRepositoryContract interface {
	GetUserByID(userId uuid.UUID) (entities.User, error)
	AddReservation(reservation entities.Reservation) error
	GetAllReservations() ([]entities.Reservation, error)
	GetReservationByID(rId uuid.UUID) (entities.Reservation, error)
	GetUserReservations(userId uuid.UUID) ([]entities.Reservation, error)
	CancelReservation(rId uuid.UUID) error
}

type TicketRepositoryContract interface {
	GetTicketByID(ticketId uuid.UUID) (entities.Ticket, error) // Not Done
	CancelTicket(ticketId uuid.UUID) error
}
