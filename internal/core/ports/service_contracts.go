package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
)

type SortFilterServiceContract interface {
	FilterFlightList(PlaneType string, t1 int, t2 int, RemainSeat string) []entities.Flight
	SortFlightList(Desc, Sortby string) []entities.Flight
}
