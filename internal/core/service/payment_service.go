package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
)

type PaymentService struct {
	redisDb ports.PaymentDbContract
	gormDb  ports.PaymentGormContract
}

func NewPaymentService(redisdb ports.PaymentDbContract, gormdb ports.PaymentGormContract) *PaymentService {
	return &PaymentService{
		redisDb: redisdb,
		gormDb:  gormdb,
	}

}

const url string = "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/services/pgw?wsdl=null"

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
	for _, passengerid := range reserve.Passengers {
		passenger, err := p.gormDb.GetPassengerById(passengerid)
		if err != nil {
			return "", errors.New("passemger not found")
		}
		passengers = append(passengers, calculateAgeByBirth(passenger.DateOfBirth))
	}

	amount := CalculatePrice(passengers, flight, reserve)
	terminalID := os.Getenv("BANK_TERMINAL_ID")
	userName := os.Getenv("BANK_USERNAME")
	userPassword := os.Getenv("BANK_USER_PASSWORD")
	localDate := time.Now().Format("20060102")
	localTime := time.Now().Format("150405")
	additionalData := ""
	callBackURL := os.Getenv("BANK_CALLBACK_URL")
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(RequestXMLBody, terminalID, userName, userPassword, id.ID(), amount, localDate, localTime, additionalData, callBackURL, payerID))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "http://interfaces.core.sw.bps.com/bpPayRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer res.Body.Close()
	print(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	var envelope EnvelopeRequest
	xml.Unmarshal(body, &envelope)
	result := strings.Split(envelope.Body.BpPayRequestResponse.Return.Text, ",")

	code := result[0]
	if code == "0" {
		refID := result[1]
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
		fmt.Println("Error:", envelope.Body.BpPayRequestResponse.Return.Text)
		return "", errors.New("invalid orderID")
	}

}

func (p *PaymentService) VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error) {
	terminalID := os.Getenv("BANK_TERMINAL_ID")
	userName := os.Getenv("BANK_USERNAME")
	userPassword := os.Getenv("BANK_USER_PASSWORD")

	if _, ok, err := p.redisDb.VerifyPaymentRequest(PayerID, RefId, orderId); err != nil || !ok {
		return ok, err
	}

	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(VerifyXMLBody, terminalID, userName, userPassword, orderId, orderId, SaleReferenceId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "http://interfaces.core.sw.bps.com/bpVerifyRequest")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var envelope EnvelopeVerify
	xml.Unmarshal(body, &envelope)

	code := strings.Split(envelope.Body.BpVerifyRequestResponse.Return.Text, ",")[0]
	if code == "0" {
		err := p.gormDb.IssueATicket(RefId)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, errors.New("invalid orderID")
	}

}
