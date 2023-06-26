package service_test

import (
	"fmt"
	"testing"

	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddReservationService(t *testing.T) {
	svc := service.NewReserveService()

	flightId, userId, pIds := uuid.New(), uuid.New(), []uuid.UUID{uuid.New(), uuid.New()}
	r, err := svc.Reserve(flightId, userId, pIds)

	assert.NoError(t, err)
	res, err := svc.GetReservationByID(r)
	assert.NoError(t, err)

	assert.Equal(res, r, res.ID)

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
