package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/go-redis/redis/v8"
)

// قرارداد های سرویس هامون

type UserServiceContract interface {
	IsUserAlreadyRegisters(newUser entities.User) bool
	AddUser(newUser entities.User) error
	LoginHandler(user entities.User) (string, error)
	GetAllUserPassengers(id string) ([]entities.Passenger, error)
	AddPassengersToUser(userId string, passenger entities.Passenger) error
	Logout(token string) error
}

type InMemoryServiceContracts interface {
	AddToken(token string)
	RevokeToken(token string) *redis.StatusCmd
	TokenReceiver() (string, error)
}
