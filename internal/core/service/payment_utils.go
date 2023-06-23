package service

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
)

type Return struct {
	Text string `xml:",chardata"`
}

type BpVerifyRequestResponse struct {
	XMLName xml.Name `xml:"bpVerifyRequestResponse"`
	Return  Return   `xml:"return"`
}

type BpPayRequestResponse struct {
	XMLName xml.Name `xml:"bpPayRequestResponse"`
	Return  Return   `xml:"return"`
}

type VerifyBody struct {
	XMLName                 xml.Name                `xml:"Body"`
	BpVerifyRequestResponse BpVerifyRequestResponse `xml:"bpVerifyRequestResponse"`
}

type RequestBody struct {
	XMLName              xml.Name             `xml:"Body"`
	BpPayRequestResponse BpPayRequestResponse `xml:"bpPayRequestResponse"`
}

type EnvelopeRequest struct {
	XMLName xml.Name    `xml:"Envelope"`
	Body    RequestBody `xml:"Body"`
}

type EnvelopeVerify struct {
	XMLName xml.Name   `xml:"Envelope"`
	Body    VerifyBody `xml:"Body"`
}

func calculateAgeByBirth(birthDate time.Time) int {
	currentDate := time.Now()
	age := currentDate.Year() - birthDate.Year()
	if currentDate.Before(time.Date(currentDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}

	return age
}

func CalculatePrice(passengersAge []int, flight entities.Flight, reserve entities.Reservation) string {
	var totalPrice float64
	for _, passengeer := range passengersAge {
		switch {
		case passengeer < 3:
			totalPrice += float64(flight.FareClass.InfantFare + flight.Tax)
		case passengeer < 18:
			totalPrice += float64(flight.FareClass.ChildFare + flight.Tax)
		default:
			totalPrice += float64(flight.FareClass.AdultFare + flight.Tax)
		}
	}
	return strconv.FormatFloat(totalPrice, 'f', -1, 64)
}

var RequestXMLBody string = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
	<soapenv:Header/>
	<soapenv:Body>
		<web:bpPayRequest>
			<web:terminalId>%s</web:terminalId>
			<web:userName>%s</web:userName>
			<web:userPassword>%s</web:userPassword>
			<web:orderId>%s</web:orderId>
			<web:amount>%s</web:amount>
			<web:localDate>%s</web:localDate>
			<web:localTime>%s</web:localTime>
			<web:additionalData>%s</web:additionalData>
			<web:callBackUrl>%s</web:callBackUrl>
			<web:payerId>%s</web:payerId>
		</web:bpPayRequest>
	</soapenv:Body>
</soapenv:Envelope>`

var VerifyXMLBody string = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
	<soapenv:Header/>
	<soapenv:Body>
		<web:bpVerifyRequest>
			<web:terminalId>%s</web:terminalId>
			<web:userName>%s</web:userName>
			<web:userPassword>%s</web:userPassword>
			<web:orderId>%s</web:orderId>
			<web:saleOrderId>%s</web:saleOrderId>
			<web:saleReferenceId>%s</web:saleReferenceId>
		</web:bpVerifyRequest>
	</soapenv:Body>
</soapenv:Envelope>`
