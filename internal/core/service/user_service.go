package service

import (
	"errors"
	"log"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/go-redis/redis/v8"
)

// خود توابع سرویس

type UserService struct {
	db    ports.UserRepositoryContracts
	redis ports.InMemoryRespositoryContracts
}

func NewUserService(db ports.UserRepositoryContracts, redis ports.InMemoryRespositoryContracts) *UserService {
	return &UserService{
		db:    db,
		redis: redis,
	}
}

func (u *UserService) IsUserAlreadyRegisters(newUser entities.User) bool {
	res := u.db.IsUserAlreadyRegisters(newUser)
	if res > 0 {
		return true
	}
	return false
}

func (u *UserService) AddUser(newUser entities.User) error {

	isUserAlreadyExist := u.IsUserAlreadyRegisters(newUser)

	if isUserAlreadyExist == true {
		err := errors.New("User already registered")
		return err
	}

	err := u.db.AddUser(newUser)
	return err
}
func (u *UserService) LoginHandler(user entities.User) (string, error) {
	email := user.Email
	password := user.Password

	foundedUser, err := u.db.LoginHandler(email)

	if err != nil {

		return "", err
	}
	if foundedUser.Password != password {
		err := errors.New("email or password mismatch")
		return "", err
	}
	token := GenerateToken(foundedUser.ID, foundedUser.Email, foundedUser.Name)
	u.redis.AddToken(token)
	return token, nil
}
func (u *UserService) GetAllUserPassengers(id string) ([]entities.Passenger, error) {
	passengers, err := u.db.GetAllUserPassengers(id)

	if err != nil {
		return nil, err
	}
	return passengers, nil

}
func (u *UserService) AddPassengersToUser(userId string, passenger entities.Passenger) error {
	err := u.db.AddPassengers(passenger)
	if err != nil {
		return err
	}
	err = u.db.AddPassengerToUser(userId, passenger.ID)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) AddToken(token string) {
	err := u.redis.AddToken(token)
	if err != nil {
		log.Fatal(err)
	}
}
func (u *UserService) Logout(token string) *redis.StatusCmd {
	err := u.redis.RevokeToken(token)
	return err
}
func (u *UserService) TokenReceiver(token string) (string, error) {
	val, err := u.redis.TokenReceiver(token)
	return val, err
}
