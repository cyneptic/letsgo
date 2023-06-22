package entities

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`                  // OrderID in other tables
	UserID      uuid.UUID    `gorm:"type:uuid;Column:user_id" json:"user_id"`         // ID in User
	FlightID    uuid.UUID    `gorm:"type:uuid;Column:flight_id" json:"flight_id"`     // ID in Flight
	Passengers  []uuid.UUID  `gorm:"type:uuid[];Column:passengers" json:"passengers"` // List of passengers
	ContactInfo *ContactInfo `gorm:"embedded" json:"contact_info"`
	CreatedAt   time.Time    `json:"created_at"`
	ModifiedAt  time.Time    `json:"modified_at"`
	DeletedAt   time.Time    `json:"deleted_at"`
	Cancelled   bool         `json:"cancelled"`
}

type ReservationRequest struct {
	FlightID   uuid.UUID   `json:"flight_id"`
	UserID     uuid.UUID   `json:"user_id"`
	Passengers []uuid.UUID `json:"passengers"`
}
