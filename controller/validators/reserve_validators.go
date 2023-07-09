package validators

import (
	"errors"

	"github.com/google/uuid"
)

func ValidateReservationParams(f, u string, p []string) error {
	if f == "" || u == "" {
		return errors.New("Please Enter Parameters.")

	}
	_, err := ValidateId(f)
	if err != nil {
		return err
	}
	_, err = ValidateId(u)
	if err != nil {
		return err
	}

	for _, v := range p {
		_, err = ValidateId(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateId(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, errors.New("Please enter ID")
	}
	return uuid.Parse(id)
}
