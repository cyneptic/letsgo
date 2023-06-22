package controllers

import (
	"github.com/cyneptic/letsgo/internal/ports"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	svc ports.TicketServiceContract
}

func NewTicketHandler(svc ports.TicketServiceContract) *TicketHandler {
	return &TicketHandler{
		svc: svc,
	}
}

func AddTicketRoutes(e *echo.Echo, svc ports.TicketServiceContract) {
	handler := NewTicketHandler(svc)
	e.DELETE("/ticket/:ticketId", handler.CancelTicket)
}

func (h *TicketHandler) CancelTicket(c echo.Context) error {
	return nil
}
