package controller

import (
	"net/http"

	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

type FlightHandlers struct {
	srv ports.SortFilterServiceContract
}

func NewFlightHandler() *FlightHandlers {
	srv := service.NewFlightService()
	return &FlightHandlers{
		srv: srv,
	}
}

func RegisterFlightRoute(e *echo.Echo) {
	handler := NewFlightHandler()
	e.GET("/flights/filter", handler.FilterFlightList)
	e.GET("/flights/sort", handler.SortFlightList)
}

func (h *FlightHandlers) FilterFlightList(c echo.Context) error {
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

	sortListed := h.srv.FilterFlightList(source, destination, departing, planeType, t1, t2, remainSeat)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}

func (h *FlightHandlers) SortFlightList(c echo.Context) error {
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
	sortListed := h.srv.SortFlightList(source, destination, departing, desc, sortby)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}
