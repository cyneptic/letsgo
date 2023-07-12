package main

import (
	"log"
	"os"

	controllers "github.com/cyneptic/letsgo/controller"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load(".env")
	e := echo.New()
	controllers.RegisterPaymentRoutes(e)
	controllers.AddFlightRoutes(e)
	controllers.AddReserveRoutes(e)
	e.Use(middleware.CustomLogger)
	// _ = repositories.NewGormDatabase()

	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}
