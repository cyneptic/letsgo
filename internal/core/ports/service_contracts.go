package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type ReserveServiceContract interface {
	CancelReservation(rId uuid.UUID) error
	GetUserReservations(userId uuid.UUID) ([]entities.Reservation, error)
	GetAllReservations() ([]entities.Reservation, error)
	GetReservationByID(rId uuid.UUID) (entities.Reservation, error)
	Reserve(flightId uuid.UUID, userId uuid.UUID, passengers []uuid.UUID) (uuid.UUID, error)
}

type TicketServiceContract interface {
	IsCancellable(ticketId uuid.UUID) (bool, error)
	CancelTicket(ticketId uuid.UUID) error
}

type PaymentServiceContract interface {
	CreateNewPayment(reservationId string, payerID int64) (string, error)
	VerifyPayment(RefId, reservationId, SaleReferenceId string) (bool, error)
}

type FlightServiceContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
	RequestFlight(id string) (entities.Flight, error)
	FilterFlightList(source string, destination string, departure string, PlaneType string, t1 int, t2 int, RemainSeat uint) []entities.Flight
	SortFlightList(source, destination, departure, Desc, Sortby string) []entities.Flight
}
