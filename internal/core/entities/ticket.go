package entities

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID              uuid.UUID    `gorm:"primaryKey;type:uuid" json:"id"` // TicketID in other tables
	FlightID        uuid.UUID    `gorm:"type:uuid;Column:flight_id" json:"flight_id"`
	UserID          uuid.UUID    `gorm:"type:uuid;Column:user_id" json:"user_id"`
	ReservationID   uuid.UUID    `gorm:"type:uuid;Column:order_id" json:"order_id"`
	PaymentID       uuid.UUID    `gorm:"type:uuid;Column:payment_id" json:"payment_id"`
	TicketNumber    string       `json:"ticket_number"`
	ReferenceNumber string       `json:"reference_number"`
	Source          string       `json:"source"`
	Destination     string       `json:"destination"`
	DepartureDate   time.Time    `json:"departure_date"`
	ArrivalDate     time.Time    `json:"arrival_date"`
	AirlineName     string       `json:"airline_name"`
	Passenger       *Passenger   `gorm:"embedded" json:"passenger"`
	ContactInfo     *ContactInfo `gorm:"embedded" json:"contact_info"`
	RefundPolicy    string       `json:"refund_policy"`
	AllowedBaggage  string       `json:"allowed_baggage"`
	Status          int          `json:"status"`
	CreatedAt       time.Time    `json:"created_at"`
	ModifiedAt      time.Time    `json:"modified_at"`
	DeletedAt       time.Time    `json:"deleted_at"`
}

type ContactInfo struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
