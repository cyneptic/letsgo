package main

import (
	"github.com/cyneptic/letsgo/infrastructure/provider"
	"github.com/labstack/echo/v4"
	"log"

	"github.com/cyneptic/letsgo/controller"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	gormDb := repositories.NewGormDatabase()
	redisDb := repositories.RedisInit()
	paymentGateway := provider.NewMellatGateway()
	srvPayment := service.NewPaymentService(redisDb, gormDb, paymentGateway)
	e := echo.New()
	controller.RegisterPaymentRoutes(e, *srvPayment)

	log.Fatal(e.Start(":8080"))
}
