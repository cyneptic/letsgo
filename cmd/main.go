package main

import (
	"log"

	"github.com/cyneptic/letsgo/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	// _ = repositories.NewGormDatabase()
	e := echo.New()
	controller.RegisterFlightRoute(e)

	log.Fatal(e.Start(":8080"))
}
