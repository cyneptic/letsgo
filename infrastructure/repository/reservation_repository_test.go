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
		Passengers: []uuid.UUID{},
		ContactInfo: &entities.ContactInfo{
			Email:       email,
			PhoneNumber: "09121234567",
		},
		CreatedAt: time.Now(),
	}
}

func TestAddReservationGorm(t *testing.T) {
	g := repositories.NewGormDatabase()

	fakeR := generateFakeReservation("TestEmail@letsgo.com")
	err := g.AddReservation(fakeR)

	assert.NoError(t, err)

	var result entities.Reservation
	q := g.DB.Where("email = ?", "TestEmail@letsgo.com").First(&result)
	assert.NoError(t, q.Error)

	assert.Equal(t, []uuid.UUID{result.FlightID, result.ID, result.UserID}, []uuid.UUID{fakeR.FlightID, fakeR.ID, fakeR.UserID})

}

func TestDeleteReservation(t *testing.T) {
	g := repositories.NewGormDatabase()
	g.DB.Where("email = ?", "TestEmail@letsgo.com").Delete(entities.Reservation{})

	var result entities.Reservation
	q := g.DB.Where("email = ?", "TestEmail@letsgo.com").First(&result)

	assert.Error(t, q.Error)
}

func TestGetAllReservations(t *testing.T) {
	g := repositories.NewGormDatabase()

	for i := 0; i < 5; i++ {
		err := g.AddReservation(generateFakeReservation("TestMultipleReservations@letsgo.com"))
		assert.NoError(t, err)
	}

	result, err := g.GetAllReservations()
	assert.NoError(t, err)

	fmt.Println(result)
}
