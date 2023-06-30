package validators

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func validateDepartingDate(departingStr string) error {
	if departingStr == "" {
		return errors.New("departing is required")
	}

	departing, err := time.Parse("2006-01-02", departingStr)
	if err != nil {
		return errors.New("invalid input date format")
	}

	now := time.Now().UTC()
	if departing.Before(now) {
		return errors.New("past date is not allowed")
	}

	return nil
}

func validateForSort(desc, sort string) error {
	if desc != "" && desc != "true" {
		return errors.New("desc is not correct")
	}

	switch sort {
	case "depDate":
	case "price":
	case "duration":
	case "":
		return nil
	default:
		return fmt.Errorf("sort by %v not exists", sort)
	}

	return nil
}

func VlidateNumberForFilter(p url.Values) (int, int, error) {

	t1, err := strconv.Atoi(p.Get("t1"))
	if err != nil {
		if p.Get("t1") == "" {
			t1 = 0
		} else {
			return 0, 0, fmt.Errorf("can't convert %v to int", t1)
		}
	}

	t2, err := strconv.Atoi(p.Get("t2"))
	if err != nil {
		if p.Get("t2") == "" {
			t2 = 0
		} else {
			return 0, 0, fmt.Errorf("can't convert %v to int", t1)
		}
	}
	return t1, t2, nil
}

func VlidateRemainSeatForFilter(p url.Values) (uint64, error) {

	re, err := strconv.ParseUint(p.Get("remainSeat"), 10, 0)
	if err != nil {
		if p.Get("remainSeat") == "" {
			re = 0
		} else {
			return 0, fmt.Errorf("can't convert %v to int", re)
		}
	}
	return re, nil
}

func ValidateListFlightParam(p url.Values) error {
	if p.Get("source") == "" {
		return errors.New("source is required")
	}

	if p.Get("destination") == "" {
		return errors.New("destination is required")
	}

	err := validateDepartingDate(p.Get("departing"))
	if err != nil {
		return err
	}

	err = validateForSort(p.Get("desc"), p.Get("sort"))

	if err != nil {
		return err
	}

	return nil
}
