package main

import (
	"log"

	"github.com/cyneptic/letsgo/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	controller.AddUserRoutes(*e)

	log.Fatal(e.Start(":8080"))
}
