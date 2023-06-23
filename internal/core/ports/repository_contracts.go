package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type PaymentDbContract interface {
	SetPaymentRequest(orderID, payerID, refID, amount string) error
	VerifyPaymentRequest(refID, orderID string) (bool, error)
}

type PaymentGormContract interface {
	GetReservationById(reservationId string) (entities.Reservation, error)
	GetFlightById(reservationId string) (entities.Flight, error)
	GetPassengerById(passengerid uuid.UUID) (entities.Passenger, error)
	CreateTempTicket(reserveObj entities.Reservation, referID string) error
	IssueATicket(refID string) error
}
