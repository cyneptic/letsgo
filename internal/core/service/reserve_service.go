package service

import (
	"errors"
	"time"

	"github.com/cyneptic/letsgo/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
)

var (
	ErrNotEnoughSeats = errors.New("Not Enough Seats Left For Reservation")
)

type ReserveService struct {
	db ports.ReserveRepositoryContract
	pv ports.ReserveProviderContract
}

func NewReserveService() *ReserveService {
	repo := repositories.NewGormDatabase()
	pv := provider.NewReservationProviderClient()
	return &ReserveService{
		db: repo,
		pv: pv,
	}
}

func (svc *ReserveService) GetReservationByID(rId uuid.UUID) (entities.Reservation, error) {
	return svc.db.GetReservationByID(rId)
}

func (svc *ReserveService) CancelReservation(rId uuid.UUID) error {
	res, err := svc.db.GetReservationByID(rId)
	if err != nil {
		return err
	}

	err = svc.pv.RequestCancelReservation(res.FlightID, len(res.Passengers))
	if err != nil {
		return err
	}

	err = svc.db.CancelReservation(rId)
	return err
}

func (svc *ReserveService) GetUserReservations(userId uuid.UUID) ([]entities.Reservation, error) {
	result, err := svc.db.GetUserReservations(userId)
	return result, err
}

func (svc *ReserveService) GetAllReservations() ([]entities.Reservation, error) {
	result, err := svc.db.GetAllReservations()
	return result, err
}

func (svc *ReserveService) Reserve(flightId uuid.UUID, userId uuid.UUID, passengers []uuid.UUID) (uuid.UUID, error) {
	// Get the flight details from provider
	flight, err := svc.pv.RequestFlightByID(flightId)
	if err != nil {
		return uuid.UUID{}, err
	}

	// check if there are enough remaining  seats
	if int(flight.RemainingSeat) < len(passengers) {
		return uuid.UUID{}, ErrNotEnoughSeats
	}

	// Get user info from repository
	user, err := svc.db.GetUserByID(userId)
	if err != nil {
		return uuid.UUID{}, err
	}

	// generate a reservation
	r := entities.Reservation{
		ID:         uuid.New(),
		UserID:     userId,
		FlightID:   flightId,
		Passengers: passengers,
		ContactInfo: &entities.ContactInfo{
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		},
		CreatedAt: time.Now(),
		Cancelled: false,
	}

	// ask provider to change the remaining seats
	err = svc.pv.RequestReserve(flightId, len(passengers))
	if err != nil {
		return uuid.UUID{}, err
	}

	// save reservation in repository
	err = svc.db.AddReservation(r)
	if err != nil {
		return uuid.UUID{}, err
	}

	// return reservation ID
	return r.ID, nil
}
