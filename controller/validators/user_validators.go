package validators

import (
	"errors"
	"reflect"
	"time"

	"github.com/cyneptic/letsgo/internal/core/entities"
)

func ValidateUserID(userId string) error {
	if userId == "" {
		return errors.New("userId cannot be empty")
	}
	return nil
}

func ValidateAuthToken(authToken string) error {
	if authToken == "" {
		return errors.New("token cannot be empty")
	}
	return nil
}

func ValidatePassenger(p entities.Passenger) error {
	if p.FirstName == "" || len(p.FirstName) < 4 {
		return errors.New("passenger FirstName should be min 4 character and cannot be empty")
	}
	if p.LastName == "" || len(p.LastName) < 4 {
		return errors.New("passenger LastName should be min 4 character and cannot be empty")
	}
	if reflect.TypeOf(p.DateOfBirth) != reflect.TypeOf(time.Time{}) {
		return errors.New("passenger DateOfBirth should be a time format")
	}
	if p.Nationality == "" {
		return errors.New("passenger Nationality cannot be empty")
	}
	if p.NationalCode == "" || len(p.NationalCode) < 10 {
		return errors.New("passenger NationalCode cannot be empty and should be min 10 character")
	}
	if p.Gender == "" {
		return errors.New("passenger Gender cannot be empty")
	}
	return nil
}

func ValidateUserLogin(u entities.User) error {
	if u.Email == "" {
		return errors.New("user email cannot be empty")
	}
	if u.Password == "" {
		return errors.New("user email cannot be empty")
	}
	return nil
}
func ValidateUserRegister(u entities.User) error {
	if u.Name == "" || len(u.Name) < 4 {
		return errors.New("user Name should be min 4 character and cannot be empty")
	}	
	if reflect.TypeOf(u.DateOfBirth) != reflect.TypeOf(time.Time{}) {
		return errors.New("user DateOfBirth should be a time format")
	}
	if u.PhoneNumber == "" || len(u.PhoneNumber) < 11 {
		return errors.New("user PhoneNumber should be at least 10 number and cannot be empty")
	}
	if u.Email == ""  {
		return errors.New("user Email cannot be empty")
	}
	if u.Password == "" || len(u.Password) < 6  {
		return errors.New("user Password must be min 6 digit and cannot be empty")
	}
	return nil
}
