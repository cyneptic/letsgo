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
	GetTicketByID(ticketId uuid.UUID) (entities.Ticket, error)
	CancelTicket(ticketId uuid.UUID) error
}

type PaymentDbContract interface {
	SetPaymentRequest(orderID uuid.UUID, payerID, refID string) error
	VerifyPaymentRequest(payerID, refID, orderID string) (string, bool, error)
}

type PaymentGormContract interface {
	GetReservationById(reservationId string) (entities.Reservation, error)
	GetFlightById(reservationId string) (entities.Flight, error)
	GetPassengerById(passengerid uuid.UUID) (entities.Passenger, error)
	CreateTempTicket(reserveObj entities.Reservation, referID string) error
	IssueATicket(refID string) error
}
