package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	IsUserAlreadyRegisters(user entities.User) int64
	AddUser(user entities.User) error
	LoginHandler(email string) (*entities.User, error)
	GetAllUserPassengers(id string) ([]entities.Passenger, error)
	AddPassengers(passenger entities.Passenger) error
	AddPassengerToUser(userId string, passengerId uuid.UUID) error
}

type InMemoryRespositoryContracts interface {
	AddToken(token string) *redis.StatusCmd
	RevokeToken(token string) *redis.StatusCmd
	TokenReceiver(token string) (string, error)
}
