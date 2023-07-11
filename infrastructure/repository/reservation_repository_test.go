package repositories_test

import (
	"fmt"
	"testing"
	"time"

	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func generateFakeReservation(email string) entities.Reservation {
	return entities.Reservation{
		ID:         uuid.New(),
		UserID:     uuid.New(),
		FlightID:   uuid.New(),
		Passengers: entities.CustomUUIDArray{uuid.New()},
		ContactInfo: &entities.ContactInfo{
			Email:       email,
			PhoneNumber: "09121234567",
		},
		CreatedAt: time.Now(),
	}
}

func generateFakeMultiPersonReservation() entities.Reservation {
	var passengers entities.CustomUUIDArray
	for i := 0; i < 5; i++ {
		passengers = append(passengers, uuid.New())
	}
	return entities.Reservation{
		ID:         uuid.New(),
		UserID:     uuid.New(),
		FlightID:   uuid.New(),
		Passengers: passengers,
		ContactInfo: &entities.ContactInfo{
			Email:       "test@test.com",
			PhoneNumber: "09121234567",
		},
		CreatedAt: time.Now(),
	}
}

func TestAddReservationForMultiplePassengers(t *testing.T) {
	g := repositories.NewPGDatabase()

	fakeR := generateFakeMultiPersonReservation()
	err := g.AddReservation(fakeR)

	assert.NoError(t, err)
}

func TestAddReservation(t *testing.T) {
	g := repositories.NewPGDatabase()

	fakeR := generateFakeReservation("TestEmail@letsgo.com")
	err := g.AddReservation(fakeR)

	assert.NoError(t, err)

	var result entities.Reservation
	q := g.DB.Where("email = ? ", "TestEmail@letsgo.com").First(&result)
	assert.NoError(t, q.Error)

	assert.Equal(t, []uuid.UUID{result.FlightID, result.ID, result.UserID}, []uuid.UUID{fakeR.FlightID, fakeR.ID, fakeR.UserID})

}

func TestDeleteReservation(t *testing.T) {
	g := repositories.NewPGDatabase()
	g.DB.Where("email = ?", "TestEmail@letsgo.com").Delete(entities.Reservation{})

	var result entities.Reservation
	q := g.DB.Where("email = ?", "TestEmail@letsgo.com").First(&result)

	assert.Error(t, q.Error)
}

func TestGetAllReservations(t *testing.T) {
	g := repositories.NewPGDatabase()

	for i := 0; i < 5; i++ {
		err := g.AddReservation(generateFakeReservation("TestMultipleReservations@letsgo.com"))
		assert.NoError(t, err)
	}

	result, err := g.GetAllReservations()
	assert.NoError(t, err)

	fmt.Println(result)
}
