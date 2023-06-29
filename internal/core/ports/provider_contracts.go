package ports

import "github.com/google/uuid"

type PaymentGateWayContract interface {
	CreatePayment(amount string, order uuid.UUID, payerID string) (status, refID string, err error)
	VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error)
}
