package main

import (
	controllers "github.com/cyneptic/letsgo/controller"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	_ = godotenv.Load(".env")
	e := echo.New()
	controllers.RegisterPaymentRoutes(e)
	controllers.AddFlightRoutes(e)
	e.Use(middleware.CustomLogger)

	log.Fatal(e.Start(":8080"))
}
