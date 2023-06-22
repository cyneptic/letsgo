package service

import (
	"errors"

	"github.com/cyneptic/letsgo/internal/ports"
	"github.com/google/uuid"
)

var (
	ErrTicketNotCancellable = errors.New("Ticket is no longer cancellable")
)

type TicketService struct {
	db ports.TicketRepositoryContract
	pv ports.TicketProviderContract
}

func NewTicketService(db ports.TicketRepositoryContract, pv ports.TicketProviderContract) *TicketService {
	return &TicketService{
		db: db,
		pv: pv,
	}
}

func (svc *TicketService) IsCancellable(ticketId uuid.UUID) (bool, error) {
	// Get Ticket (set to _ for now)
	_, err := svc.db.GetTicketByID(ticketId)
	if err != nil {
		return false, err
	}

	// How to do this?!

	return true, nil
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
