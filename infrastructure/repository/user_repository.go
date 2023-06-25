package repositories

import (
	"context"
	// "fmt"
	"time"

	// "strings"

	"github.com/cyneptic/letsgo/internal/core/entities"
	
	
)

func (p *Postgres) IsUserAlreadyRegisters(user entities.User) int64 {
	res := p.db.Where("email = ?", user.Email).First(&user)
	return res.RowsAffected
}
func (p *Postgres) AddUser(user entities.User) error {
	result := p.db.Create(user)
	return result.Error
}
func (p *Postgres) LoginHandler(email string) (*entities.User, error) {

	var fundedUser entities.User
	if err := p.db.Where("email = ? ", email).First(&fundedUser).Error; err != nil {
		return nil, err
	}
	return &fundedUser, nil
}
func (p *Postgres) GetAllUserPassengers(id string) ([]entities.Passenger, error) {
	var passengers []entities.Passenger

	if err := p.db.Where("user_id = ?", id).Find(&passengers).Error; err != nil {
		return nil, err
	}
	return passengers, nil
}
func (p *Postgres) AddPassengers(passenger entities.Passenger) error {
	result := p.db.Create(passenger)
	return result.Error
}
func (p *Postgres) AddPassengerToUser(userId string, passenger entities.Passenger) error {
	user := entities.User{}
	if err := p.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	user.Passengers = append(user.Passengers, passenger)

	if err := p.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// redis
func (r *RedisDB) AddToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.client.Set(ctx, token, true, 0).Err()
	if err != nil {
        panic(err)
    }
	return err

}
func (r *RedisDB) RevokeToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.client.Set(ctx, token, false, 0).Err()
	return err
}

func (r *RedisDB) TokenReceiver(token string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, token).Result()

	return val, err
}
