package ports

import "github.com/cyneptic/letsgo/internal/core/entities"

type PaymentServiceContract interface {
	CreateNewPayment(reservationId string, payerID int64) (string, error)
	VerifyPayment(RefId, reservationId, SaleReferenceId string) (bool, error)
}

type FlightServiceContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
	RequestFlight(id string) (entities.Flight, error)
}
