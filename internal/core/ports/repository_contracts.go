package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	

)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	IsUserAlreadyRegisters(user entities.User) (int64 , error)
	AddUser(user entities.User) error
	LoginHandler(email string) (*entities.User, error)
	GetAllUserPassengers(id string) ([]entities.Passenger, error)
	AddPassengers(passenger entities.Passenger) error
	AddPassengerToUser(userId string, passengerId entities.Passenger) error
}

type InMemoryRespositoryContracts interface {
	AddToken(token string) error
	RevokeToken(token string) error
	TokenReceiver(token string) (string, error)
}
