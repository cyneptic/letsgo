package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReservationRequest struct {
	FlightID   uuid.UUID   `json:"flight_id"`
	UserID     uuid.UUID   `json:"user_id"`
	Passengers []uuid.UUID `json:"passengers"`
}

type ReservationHandler struct {
	svc ports.ReserveServiceContract
}

func NewReserveHandler() *ReservationHandler {
	svc := service.NewReserveService()
	return &ReservationHandler{
		svc: svc,
	}
}

func AddReserveRoutes(e *echo.Echo) {
	handler := NewReserveHandler()
	e.POST("/reserve", handler.Reserve)
	e.GET("/reserve", handler.AllReservations)
	e.GET("/reserve/user/:user_id", handler.UserReservations)
	e.GET("/reserve/:reservation_id", handler.GetReservationByID)
	e.DELETE("/reserve/:reservation_id", handler.Cancel)
}

func (h *ReservationHandler) GetReservationByID(c echo.Context) error {
	rId := c.QueryParam("reservation_id")
	parsedId, err := validators.ValidateId(rId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	r, err := h.svc.GetReservationByID(parsedId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, r)

}

func (h *ReservationHandler) Reserve(c echo.Context) error {
	var reservationForm ReservationRequest
	if err := c.Bind(&reservationForm); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	rId, err := h.svc.Reserve(reservationForm.FlightID, reservationForm.UserID, reservationForm.Passengers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, rId)
}

func (h *ReservationHandler) AllReservations(c echo.Context) error {
	reservations, err := h.svc.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) UserReservations(c echo.Context) error {
	uId := c.QueryParam("user_id")
	userId, err := validators.ValidateId(uId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reservations, err := h.svc.GetUserReservations(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) Cancel(c echo.Context) error {
	rId := c.QueryParam("reservation_id")
	resId, err := validators.ValidateId(rId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.svc.CancelReservation(resId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}
