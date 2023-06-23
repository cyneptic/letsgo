package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CustomUUIDArray []uuid.UUID

// Value returns the database value of CustomUUIDArray
func (a CustomUUIDArray) Value() (driver.Value, error) {
	var uuidStrings []string
	for _, u := range a {
		uuidStrings = append(uuidStrings, u.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(uuidStrings, ",")), nil
}

// Scan scans the database value into CustomUUIDArray
func (a *CustomUUIDArray) Scan(value interface{}) error {
	if value == nil {
		*a = CustomUUIDArray{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var uuidStrings []string
		err := json.Unmarshal(v, &uuidStrings)
		if err != nil {
			return err
		}

		var uuids []uuid.UUID
		for _, s := range uuidStrings {
			uuid, err := uuid.Parse(s)
			if err != nil {
				return err
			}
			uuids = append(uuids, uuid)
		}

		*a = CustomUUIDArray(uuids)
	case string:
		if len(v) < 2 {
			*a = CustomUUIDArray{}
			return nil
		}

		v = v[1 : len(v)-1] // حذف علامت‌های { و }
		uuidStrings := strings.Split(v, ",")
		var uuids []uuid.UUID
		for _, s := range uuidStrings {
			uuid, err := uuid.Parse(strings.TrimSpace(s))
			if err != nil {
				return err
			}
			uuids = append(uuids, uuid)
		}

		*a = CustomUUIDArray(uuids)
	default:
		return fmt.Errorf("unsupported Scan type for CustomUUIDArray: %T", value)
	}

	return nil
}

type Reservation struct {
	ID             uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`                 // OrderID in other tables
	UserID         uuid.UUID       `gorm:"type:uuid;Column:user_id" json:"user_id"`         // ID in User
	FlightID       uuid.UUID       `gorm:"type:uuid;Column:flight_id" json:"flight_id"`     // ID in Flight
	Passengers     CustomUUIDArray `gorm:"type:uuid[];Column:passengers" json:"passengers"` // List of passengers
	ContactInfo    *ContactInfo    `gorm:"embedded" json:"contact_info"`
	Source         string          `json:"source"`
	Destination    string          `json:"destination"`
	DepartureDate  time.Time       `json:"departure_date"`
	ArrivalDate    time.Time       `json:"arrival_date"`
	AirlineName    string          `json:"airline_name"`
	RefundPolicy   string          `json:"refund_policy"`
	AllowedBaggage string          `json:"allowed_baggage"`
	CreatedAt      time.Time       `json:"created_at"`
	ModifiedAt     time.Time       `json:"modified_at"`
	DeletedAt      time.Time       `json:"deleted_at"`
	Cancelled      bool            `json:"cancelled"`
}

type ReservationRequest struct {
	FlightID   uuid.UUID   `json:"flight_id"`
	UserID     uuid.UUID   `json:"user_id"`
	Passengers []uuid.UUID `json:"passengers"`
}
