package entities

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"` // FlightID in other tables
	FlightNumber   string    `json:"flight_number"`                  // FlightNumber - 12587129381295
	Source         string    `json:"source"`                         // Source Airport
	Destination    string    `json:"destination"`                    // Destination Airpot
	DepartureDate  time.Time `json:"departure_date"`                 // Departure Time and Date
	FlightDuration int       `json:"flight_duration"`
	ArrivalDate    time.Time `json:"arrival_date"` // Arrival Time and Date
	AirlineName    string    `json:"airline_name"` // Mahan - IranAir - Aseman - Homa
	// AirlineCode string
	AircraftName  string    `json:"aircraft_name"`  // Boeing-737
	FareClass     FareClass `gorm:"embedded"`       // Prices relative to age
	Tax           int64     `json:"tax"`            // Total Price == BaseFare + Tax
	FlightClass   string    `json:"flight_class"`   // Eco - Business - Eco+ - FirstClass
	RemainingSeat uint      `json:"remaining_seat"` // 20. Number of remaining seats
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type FareClass struct {
	AdultFare  int64 `json:"adult_fare"`
	ChildFare  int64 `json:"child_fare"`
	InfantFare int64 `json:"infant_fare"`
}
