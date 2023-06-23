package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

func (r *RedisDB) SetPaymentRequest(orderID, payerID, refID, amount string) error {
	key := fmt.Sprintf("%s-%s-%s", payerID, orderID, refID)
	r.client.Set(context.Background(), key, amount, 10*time.Minute)
	return nil
}
func (r *RedisDB) VerifyPaymentRequest(refID, orderID string) (bool, error) {
	keys, _ := r.client.Keys(context.TODO(), fmt.Sprintf("*-%s-%s", orderID, refID)).Result()
	if len(keys) > 0 {
		return true, nil
	}
	return false, errors.New("invalid payment")
}

func (p *PGRepository) GetReservationById(reservationId string) (entities.Reservation, error) {
	var reserve entities.Reservation

	err := p.DB.Model(&entities.Reservation{}).Find(&reserve, "id = ?", reservationId).Error
	if err != nil {
		return entities.Reservation{}, errors.New("reservation not found")
	}
	return reserve, nil
}
func (p *PGRepository) GetFlightById(flightId string) (entities.Flight, error) {
	var flight entities.Flight

	err := p.DB.Model(&entities.Flight{}).Find(&flight, "id = ?", flightId).Error
	if err != nil {
		return entities.Flight{}, errors.New("reservation not found")
	}
	return flight, nil
}

func (p *PGRepository) GetPassengerById(passengerid uuid.UUID) (entities.Passenger, error) {
	var passengeer entities.Passenger

	err := p.DB.Model(&entities.Passenger{}).Find(&passengeer, "id = ?", passengerid).Error
	if err != nil {
		return entities.Passenger{}, errors.New("reservation not found")
	}
	return passengeer, nil
}

func (p *PGRepository) CreateTempTicket(reserveObj entities.Reservation, referID string) error {
	var ticket entities.Ticket
	ticket.ID = uuid.New()
	ticket.FlightID = reserveObj.FlightID
	ticket.UserID = reserveObj.UserID
	ticket.ReservationID = reserveObj.ID
	ticket.TicketNumber = ""
	ticket.ReferenceNumber = referID
	ticket.Source = reserveObj.Source
	ticket.Destination = reserveObj.Destination
	ticket.DepartureDate = reserveObj.DepartureDate
	ticket.ArrivalDate = reserveObj.ArrivalDate
	ticket.AirlineName = reserveObj.AirlineName
	ticket.Passenger = reserveObj.Passenger
	ticket.ContactInfo = reserveObj.ContactInfo
	ticket.RefundPolicy = reserveObj.RefundPolicy
	ticket.AllowedBaggage = reserveObj.AllowedBaggage
	ticket.Status = 0
	ticket.CreatedAt = time.Now()
	ticket.ModifiedAt = time.Now()

	result := p.DB.Create(&ticket)
	if result.Error != nil {
		return errors.New("error in create object")
	}
	return nil
}

func (p *PGRepository) IssueATicket(refID string) error {
	var ticket entities.Ticket
	err := p.DB.Model(&entities.Ticket{}).Find(&ticket, "reference_number = ?", refID).Error
	if err != nil {
		return errors.New("object not found")

	}
	ticket.Status = 1

	if p.DB.Save(ticket).Error != nil {
		return errors.New("error in issuing ticket")
	}
	return nil
}
