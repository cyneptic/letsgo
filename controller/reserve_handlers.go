package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/domain/entities"
	"github.com/cyneptic/letsgo/internal/ports"
	"github.com/docker/distribution/uuid"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	svc ports.ReserveServiceContract
}

func NewReserveHasndler(svc ports.ReserveServiceContract) *ReservationHandler {
	return &ReservationHandler{
		svc: svc,
	}
}

func AddReserveRoutes(e *echo.Echo, svc ports.ReserveServiceContract) {
	handler := NewReserveHasndler(svc)
	e.POST("/reserve", handler.Reserve)
	e.GET("/reserve", handler.AllReservations)
	e.GET("/reserve/:user_id", handler.UserReservations)
	e.DELETE("/reserve/:reservation_id", handler.Cancel)
}

func (h *ReservationHandler) Reserve(c echo.Context) error {
	reservationForm := new(entities.ReservationRequest)
	if err := c.Bind(&reservationForm); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := validators.ValidateReservationParams(reservationForm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	rId, err := h.svc.Reserve(reservationForm.FlightID, reservationForm.UserID, reservationForm.Passengers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, rId)
}

func (h *ReservationHandler) AllReservations(c echo.Context) error {
	reservations, err := h.svc.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) UserReservations(c echo.Context) error {
	err := validators.ValidateUserId(c.QueryParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	reservations, err := h.svc.GetUserReservations(uuid.Parse(c.QueryParam("user_id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) Cancel(c echo.Context) error {
	rId := c.QueryParam("reservation_id")
	err := validators.ValidateReservationId(rId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err = h.svc.CancelReservation(uuid.Parse(rId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
