package controller

import (
	"net/http"
	"strconv"

	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

type FlightHandlers struct {
	srv ports.SortFilterServiceContract
}

func NewFlightHandler(srv ports.SortFilterServiceContract) *FlightHandlers {
	return &FlightHandlers{
		srv: srv,
	}
}

func RegisterFlightRoute(e *echo.Echo) {
	handler := NewFlightHandler(service.NewFlightService())
	e.GET("/flights/filter", handler.FilterFlightList)
	e.GET("/flights/sort", handler.SortFlightList)
}

func (h *FlightHandlers) FilterFlightList(c echo.Context) error {
	//! declare parameters
	var planeType = c.QueryParam("planeType")
	var t1, _ = strconv.Atoi(c.QueryParam("t1"))
	var t2, _ = strconv.Atoi(c.QueryParam("t2"))
	var remainSeat = c.QueryParam("remainSeat")

	sortListed := h.srv.FilterFlightList(planeType, t1, t2, remainSeat)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}

func (h *FlightHandlers) SortFlightList(c echo.Context) error {
	//! declare parameters
	var desc = c.QueryParam("desc")
	var sortby = c.QueryParam("sort")
	sortListed := h.srv.SortFlightList(desc, sortby)

	//! return json
	return c.JSON(http.StatusOK, sortListed)
}
