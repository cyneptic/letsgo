package repositories

import (
	"fmt"
	"os"

	"github.com/cyneptic/letsgo/internal/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostGres struct {
	db *gorm.DB
}

func GormInit() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(&entities.Passenger{}, &entities.Ticket{}, &entities.User{}, &entities.Flight{}, &entities.Reservation{})
	if err != nil {
		fmt.Println(err)
	}
	return database, nil
}
