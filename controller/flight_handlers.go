package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/infrastructure/provider"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

type FlightHandler struct {
	svc ports.FlightServiceContract
}

func NewFlightHandler() *FlightHandler {
	pv := provider.NewFlightProviderClient()
	svc := service.NewFlightService(pv)
	return &FlightHandler{
		svc: svc,
	}
}

func AddFlightRoutes(e *echo.Echo) {
	handler := NewFlightHandler()
	e.GET("/flights", handler.ListFlightsHandler)
	e.GET("/flights/:id", handler.flightHandler)
}

func (h *FlightHandler) ListFlightsHandler(c echo.Context) error {
	var flightList []entities.Flight
	err := validators.ValidateListFlightParam(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	flightList, err = h.svc.RequestFlights(c.QueryParam("source"), c.QueryParam("destination"), c.QueryParam("departing"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, flightList)
}

func (h *FlightHandler) flightHandler(c echo.Context) error {
	var flight entities.Flight
	flight, err := h.svc.RequestFlight(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, flight)
}
