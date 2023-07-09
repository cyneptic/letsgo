package service

import (
	"errors"
	"time"

	"github.com/cyneptic/letsgo/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
)

var (
	ErrTicketNotCancellable = errors.New("Ticket is no longer cancellable")
)

type TicketService struct {
	db ports.TicketRepositoryContract
	pv ports.TicketProviderContract
}

func NewTicketService() *TicketService {
	db := repositories.NewPGDatabase()
	pv := provider.NewTicketProviderClient()
	return &TicketService{
		db: db,
		pv: pv,
	}
}

func (svc *TicketService) IsCancellable(ticketId uuid.UUID) (bool, error) {
	ticket, err := svc.db.GetTicketByID(ticketId)
	if err != nil {
		return false, err
	}

	tempTime := ticket.DepartureDate.Add(10 * time.Hour) // hardcoded 10 hours, because no policy was given
	return time.Now().Before(tempTime), nil
}

func (svc *TicketService) CancelTicket(ticketId uuid.UUID) error {
	// if not cancellable return error
	cancellable, err := svc.IsCancellable(ticketId)
	if err != nil {
		return err
	}
	if !cancellable {
		return ErrTicketNotCancellable
	}

	// get ticket information from database
	ticket, err := svc.db.GetTicketByID(ticketId)
	if err != nil {
		return err
	}

	// request ticket cancellation to provider
	err = svc.pv.RequestCancelTicket(ticket.FlightID)
	if err != nil {
		return err
	}

	// cancel ticket in database
	err = svc.db.CancelTicket(ticketId)
	return err
}
