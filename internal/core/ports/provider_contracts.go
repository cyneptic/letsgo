package ports
import "github.com/cyneptic/letsgo/internal/core/entities"
import "github.com/google/uuid"

type PaymentGateWayContract interface {
	CreatePayment(amount string, order uuid.UUID, payerID string) (status, refID string, err error)
	VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error)
}

type FlightProviderContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
	RequestFlight(id string) (entities.Flight, error)
}
