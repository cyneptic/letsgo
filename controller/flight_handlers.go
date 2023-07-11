package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

type FlightHandler struct {
	svc ports.FlightServiceContract
}

func NewFlightHandler() *FlightHandler {
	svc := service.NewFlightService()
	return &FlightHandler{
		svc: svc,
	}
}

func AddFlightRoutes(e *echo.Echo) {
	handler := NewFlightHandler()
	e.GET("/flights", handler.ListFlightsHandler)
	e.GET("/flights/:id", handler.flightHandler)
	e.GET("/flights/filter", handler.FilterFlightList)
	e.GET("/flights/sort", handler.SortFlightList)
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

func (h *FlightHandler) FilterFlightList(c echo.Context) error {
	//! declare parameters
	err := validators.ValidateListFlightParam(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	remainSeat, err := validators.VlidateRemainSeatForFilter(c.QueryParams())

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	t1, t2, err := validators.VlidateNumberForFilter(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var source = c.QueryParam("source")
	var destination = c.QueryParam("destination")
	var departing = c.QueryParam("departing")
	var planeType = c.QueryParam("planeType")

	sortListed := h.svc.FilterFlightList(source, destination, departing, planeType, t1, t2, remainSeat)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}

func (h *FlightHandler) SortFlightList(c echo.Context) error {
	err := validators.ValidateListFlightParam(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//! declare parameters
	var source = c.QueryParam("source")
	var destination = c.QueryParam("destination")
	var departing = c.QueryParam("departing")
	var desc = c.QueryParam("desc")
	var sortby = c.QueryParam("sort")
	sortListed := h.svc.SortFlightList(source, destination, departing, desc, sortby)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}
