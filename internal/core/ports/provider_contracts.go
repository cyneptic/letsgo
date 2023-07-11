package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type PaymentGateWayContract interface {
	CreatePayment(amount string, order uuid.UUID, payerID string) (string, string, error)
	VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error)
}

type FlightProviderContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
	RequestFlight(id string) (entities.Flight, error)
}

type ReserveProviderContract interface {
	RequestFlightByID(flightId uuid.UUID) (entities.Flight, error)
	RequestReserve(flightId uuid.UUID, count int) error
	RequestCancelReservation(flightId uuid.UUID, count int) error
}

type TicketProviderContract interface {
	RequestCancelTicket(flightId uuid.UUID) error
}
