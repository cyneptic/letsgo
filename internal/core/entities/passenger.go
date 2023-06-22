package entities

import (
	"time"

	"github.com/google/uuid"
)

type Passenger struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;Column:user_id" json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Nationality  string    `json:"nationality"`
	NationalCode string    `json:"national_code"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
