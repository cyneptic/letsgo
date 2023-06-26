package main

import (
	"log"

	controllers "github.com/cyneptic/letsgo/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controllers.AddReserveRoutes(e)

	log.Fatal(e.Start(":8080"))
}
