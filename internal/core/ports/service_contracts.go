package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
)

type SortFilterServiceContract interface {
	FilterFlightList(source string, destination string, departure string, PlaneType string, t1 int, t2 int, RemainSeat uint64) []entities.Flight
	SortFlightList(source, destination, departure, Desc, Sortby string) []entities.Flight
}
