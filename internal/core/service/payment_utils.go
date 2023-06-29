package service

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"strconv"
	"time"
)

func calculateAgeByBirth(birthDate time.Time) int {
	currentDate := time.Now()
	age := currentDate.Year() - birthDate.Year()
	if currentDate.Before(time.Date(currentDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}

	return age
}

func CalculatePrice(passengersAge []int, flight entities.Flight, reserve entities.Reservation) string {
	var totalPrice float64
	for _, passengeer := range passengersAge {
		switch {
		case passengeer < 3:
			totalPrice += float64(flight.FareClass.InfantFare + flight.Tax)
		case passengeer < 18:
			totalPrice += float64(flight.FareClass.ChildFare + flight.Tax)
		default:
			totalPrice += float64(flight.FareClass.AdultFare + flight.Tax)
		}
	}
	return strconv.FormatFloat(totalPrice, 'f', -1, 64)
}
