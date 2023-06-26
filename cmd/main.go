package main

import (
	controllers "github.com/cyneptic/letsgo/controller"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"log"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = repositories.NewGormDatabase()
	e := echo.New()
  
	controllers.AddFlightRoutes(e)
	e.Use(middleware.CustomLogger)

	log.Fatal(e.Start(":8080"))
}
