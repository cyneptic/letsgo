package service

import (
	"errors"
	"fmt"
	"github.com/cyneptic/letsgo/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
)

type PaymentService struct {
	redisDb ports.PaymentDbContract
	gormDb  ports.PaymentGormContract
	gateway ports.PaymentGateWayContract
}

func NewPaymentService() *PaymentService {
	gormDb := repositories.NewGormDatabase()
	redisDb := repositories.RedisInit()
	paymentGateway := provider.NewMellatGateway()
	return &PaymentService{
		redisDb: redisDb,
		gormDb:  gormDb,
		gateway: paymentGateway,
	}

}

func (p *PaymentService) CreateNewPayment(reservationId, payerID string) (string, error) {
	id, _ := uuid.Parse(reservationId)
	reserve, err := p.gormDb.GetReservationById(reservationId)

	if err != nil {
		return "", err
	}

	flight, err := p.gormDb.GetFlightById(reserve.FlightID.String())
	if err != nil {
		return "", err
	}

	var passengers []int
	for _, passengerId := range reserve.Passengers {
		passenger, err := p.gormDb.GetPassengerById(passengerId)
		if err != nil {
			return "", errors.New("passenger not found")
		}
		passengers = append(passengers, calculateAgeByBirth(passenger.DateOfBirth))
	}
	amount := CalculatePrice(passengers, flight, reserve)
	code, refID, err := p.gateway.CreatePayment(amount, id, payerID)
	if code == SUCCESS_STATUS_CODE {
		redirectLink := fmt.Sprintf(`<form name="myform" action="https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/startpay.mellat" method="POST">
		<input type="hidden" id="RefId" name="RefId" value="%s">
		</form>
		<script type="text/javascript">window.onload = formSubmit; function formSubmit() { document.forms[0].submit(); }</script>
		`, refID)
		p.redisDb.SetPaymentRequest(id, payerID, refID)
		err := p.gormDb.CreateTempTicket(reserve, refID)
		if err != nil {
			return "", err
		}
		return redirectLink, nil
	} else {
		return "", errors.New("invalid orderID")
	}

}

func (p *PaymentService) VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error) {
	if _, ok, err := p.redisDb.VerifyPaymentRequest(PayerID, RefId, orderId); err != nil || !ok {
		return ok, err
	}
	ok, _ := p.gateway.VerifyPayment(PayerID, RefId, orderId, SaleReferenceId)
	if ok {
		err := p.gormDb.IssueATicket(RefId)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, errors.New("invalid orderID")
	}

}
