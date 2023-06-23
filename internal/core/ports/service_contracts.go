package ports

type PaymentServiceContract interface {
	CreateNewPayment(reservationId string, payerID int64) (string, error)
	VerifyPayment(RefId, reservationId, SaleReferenceId string) (bool, error)
}
