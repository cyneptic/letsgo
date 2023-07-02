package service

import (
	"sort"

	"github.com/cyneptic/letsgo/infrastructure/provider"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
)

type FlightService struct {
	pv ports.FlightProviderContract
}

func NewFlightService() *FlightService {
	pv := provider.NewFlightProviderClient()
	return &FlightService{
		pv: pv,
	}
}

func (svc *FlightService) RequestFlights(source, destination, departure string) ([]entities.Flight, error) {
	return svc.pv.RequestFlights(source, destination, departure)
}
func (svc *FlightService) RequestFlight(id string) (entities.Flight, error) {
	return svc.pv.RequestFlight(id)
}

// Todo Filter
func (svc *FlightService) FilterFlightList(source string, destination string, departure string, PlaneType string, t1 int, t2 int, RemainSeat uint) []entities.Flight {
	flights, err := svc.pv.RequestFlights(source, destination, departure)
	if err != nil {
		panic(err)
	}

	var filteredListed []entities.Flight
	var remainFligtlist []entities.Flight
	var hourFilterFlight []entities.Flight

	//?filter
	for _, flight := range flights {
		if flight.RemainingSeat > RemainSeat {
			remainFligtlist = append(remainFligtlist, flight)
		}

	}

	for _, flight := range remainFligtlist {
		if flight.AirlineName == PlaneType || PlaneType == "" {
			filteredListed = append(filteredListed, flight)
		}
	}
	for _, flight := range filteredListed {
		hour := flight.DepartureDate.Hour()
		if (t1 == 0 && t2 == 0) || t1 == 0 {
			return filteredListed
		} else if t2 == 0 {
			if hour >= t1 && hour <= t1+1 {
				hourFilterFlight = append(hourFilterFlight, flight)
			}
		} else if hour >= t1 && hour <= t2 {
			hourFilterFlight = append(hourFilterFlight, flight)
		}
	}
	return hourFilterFlight
}

// todo Sort
func (svc *FlightService) SortFlightList(source, destination, departure, Desc, Sortby string) []entities.Flight {
	flights, err := svc.pv.RequestFlights(source, destination, departure)
	if err != nil {
		panic(err)
	}
	//?Sort
	sortListed := flights
	// //!Logic
	if Sortby == "depDate" {
		sort.Slice(sortListed, func(i, j int) bool {

			// t1, _ := time.Parse(time.RFC3339, sortListed[i].DepartureDate)
			// t2, _ := time.Parse(time.RFC3339, sortListed[j].DepartureDate)

			return sortListed[i].DepartureDate.Unix() < sortListed[j].DepartureDate.Unix()
		})
	}

	if Sortby == "price" {
		sort.Slice(sortListed, func(i, j int) bool {
			if sortListed[i].FareClass.AdultFare != sortListed[j].FareClass.AdultFare {
				return sortListed[i].FareClass.AdultFare < sortListed[j].FareClass.AdultFare
			} else if sortListed[i].FareClass.ChildFare != sortListed[j].FareClass.ChildFare {
				return sortListed[i].FareClass.ChildFare < sortListed[j].FareClass.ChildFare
			} else if sortListed[i].FareClass.InfantFare != sortListed[j].FareClass.InfantFare {
				return sortListed[i].FareClass.InfantFare < sortListed[j].FareClass.InfantFare
			} else {
				return sortListed[i].Tax < sortListed[j].Tax
			}

		})
	}

	if Sortby == "duration" {
		sort.Slice(sortListed, func(i, j int) bool {
			// td1, _ := time.Parse(time.TimeOnly, sortListed[i].DepartureTime)
			// td2, _ := time.Parse(time.TimeOnly, sortListed[j].DepartureTime)

			// ta1, _ := time.Parse(time.TimeOnly, sortListed[i].ArrivalTime)
			// ta2, _ := time.Parse(time.TimeOnly, sortListed[j].ArrivalTime)

			// t1 := td1.Sub(ta1)
			// t2 := td2.Sub(ta2)

			return sortListed[i].FlightDuration < sortListed[j].FlightDuration
		})
	}

	if Desc == "true" {
		for i, j := 0, len(sortListed)-1; i < j; i, j = i+1, j-1 {
			sortListed[i], sortListed[j] = sortListed[j], sortListed[i]
		}
	}
	return sortListed
}
