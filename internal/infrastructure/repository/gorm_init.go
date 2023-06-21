package repositories

import (
	"fmt"

	"github.com/cyneptic/letsgo/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInit() (*gorm.DB, error) {
	host := "localhost"      // Ideal situation this would go in a env file
	user := "postgres"       // Ideal situation this would go in a env file
	password := "mypassword" // Ideal situation this would go in a env file
	dbName := "postgres"     // Ideal situation this would go in a env file
	port := 5432             // Ideal situation this would go in a env file

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
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
