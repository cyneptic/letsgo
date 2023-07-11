package validators

import (
	"errors"
	"net/url"
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

	return nil
}
