package main

import (
	"log"

	"github.com/cyneptic/letsgo/controller"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load("../.env")
	gormDb := repositories.NewGormDatabase()
	redisDb := repositories.RedisInit()
	uuid.Parse("4324902342384")

	srvPayment := service.NewPaymentService(redisDb, gormDb)
	e := echo.New()
	controller.RegisterPaymentRoutes(e, *srvPayment)

	log.Fatal(e.Start(":8080"))
}
