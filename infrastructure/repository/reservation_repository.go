package repositories

import (
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PGRepository struct {
	DB *gorm.DB
}

func NewGormDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}

func (r *PGRepository) CancelTicket(ticketId uuid.UUID) error {
	var res entities.Ticket
	q := r.DB.Model(&res).Where("id = ?", ticketId).Update("Status", "Cancelled")
	if q.Error != nil {
		return q.Error
	}
	q = r.DB.Model(&res).Where("id = ?", ticketId).Update("Status", "Cancelled")
	return q.Error
}

func (r *PGRepository) CancelReservation(rId uuid.UUID) error {
	var res entities.Reservation
	q := r.DB.Model(&res).Where("id = ?", rId).Update("Cancelled", true)
	if q.Error != nil {
		return q.Error
	}
	q = r.DB.Model(&res).Where("id = ?", rId).Update("modified_at", time.Now())
	return q.Error
}

func (r *PGRepository) GetReservationByID(rId uuid.UUID) (entities.Reservation, error) {
	var result entities.Reservation
	q := r.DB.Where("id = ?", rId).First(&result)
	return result, q.Error
}

func (r *PGRepository) GetUserReservations(userId uuid.UUID) ([]entities.Reservation, error) {
	var result []entities.Reservation
	q := r.DB.Where("user_id = ?", userId).Find(&result)
	return result, q.Error
}

// Return All Reservations in the database or return an error
func (r *PGRepository) GetAllReservations() ([]entities.Reservation, error) {
	var result []entities.Reservation
	q := r.DB.Find(&result)
	return result, q.Error
}

// insert a reservation into database
func (r *PGRepository) AddReservation(reservation entities.Reservation) error {
	q := r.DB.Save(reservation)
	if q.Error != nil {
		return q.Error
	}
	return nil
}
