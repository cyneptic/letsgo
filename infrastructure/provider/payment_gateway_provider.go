package provider

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type MellatGateway struct {
	url string
}

func NewMellatGateway() *MellatGateway {
	return &MellatGateway{
		url: "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/services/pgw?wsdl=null",
	}
}

func (m *MellatGateway) CreatePayment(amount string, order uuid.UUID, payerID string) (status, refID string, err error) {

	terminalID := os.Getenv("BANK_TERMINAL_ID")
	userName := os.Getenv("BANK_USERNAME")
	userPassword := os.Getenv("BANK_USER_PASSWORD")
	localDate := time.Now().Format("20060102")
	localTime := time.Now().Format("150405")
	additionalData := ""
	callBackURL := os.Getenv("BANK_CALLBACK_URL")
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(service.RequestXMLBody, terminalID, userName, userPassword, order.ID(), amount, localDate, localTime, additionalData, callBackURL, payerID))

	client := &http.Client{}
	req, err := http.NewRequest(method, m.url, payload)

	if err != nil {
		return "", "", err
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "http://interfaces.core.sw.bps.com/bpPayRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	defer res.Body.Close()
	print(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	var envelope service.EnvelopeRequest
	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		return "", "", errors.New("there is error in marshaling")
	}
	result := strings.Split(envelope.Body.BpPayRequestResponse.Return.Text, ",")
	if len(result) > 1 {

		return result[0], result[1], nil
	}
	return result[0], "", nil
}

func (m *MellatGateway) VerifyPayment(PayerID, RefId, orderId, SaleReferenceId string) (bool, error) {
	terminalID := os.Getenv("BANK_TERMINAL_ID")
	userName := os.Getenv("BANK_USERNAME")
	userPassword := os.Getenv("BANK_USER_PASSWORD")
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(service.VerifyXMLBody, terminalID, userName, userPassword, orderId, orderId, SaleReferenceId))

	client := &http.Client{}
	req, err := http.NewRequest(method, m.url, payload)

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

	var envelope service.EnvelopeVerify
	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		return false, err
	}
	code := strings.Split(envelope.Body.BpVerifyRequestResponse.Return.Text, ",")[0]
	return code == service.SUCCESS_STATUS_CODE, nil
}
