package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	svc ports.TicketServiceContract
}

func NewTicketHandler() *TicketHandler {
	svc := service.NewTicketService()
	return &TicketHandler{
		svc: svc,
	}
}

func AddTicketRoutes(e *echo.Echo, svc ports.TicketServiceContract) {
	handler := NewTicketHandler()
	e.DELETE("/ticket/:ticketId", handler.CancelTicket)
}

func (h *TicketHandler) CancelTicket(c echo.Context) error {
	tId := c.QueryParam("ticket_id")
	ticketId, err := validators.ValidateId(tId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.svc.CancelTicket(ticketId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
