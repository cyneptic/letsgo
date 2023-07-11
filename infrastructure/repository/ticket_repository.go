package repositories

import (
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

func (r *PGRepository) GetTicketByID(ticketId uuid.UUID) (entities.Ticket, error) {
	var res entities.Ticket
	if err := r.DB.Model(&entities.Ticket{}).Where("id = ?", ticketId).First(&res).Error; err != nil {
		return entities.Ticket{}, err
	}

	return res, nil
}

func (r *PGRepository) CancelTicket(ticketId uuid.UUID) error {
	var res entities.Ticket
	q := r.DB.Model(&res).Where("id = ?", ticketId).Update("Status", "Cancelled")
	if q.Error != nil {
		return q.Error
	}

	q = r.DB.Model(&res).Where("id = ?", ticketId).Update("Status", "Cancelled")
	if q.Error != nil {
		return q.Error
	}

	q = r.DB.Model(&res).Where("id = ?", ticketId).Update("modified_at", time.Now())
	return q.Error
}
