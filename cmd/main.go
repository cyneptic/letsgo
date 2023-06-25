package main

import (
	"log"

	"github.com/cyneptic/letsgo/controller/middleware"
	repositories "github.com/cyneptic/letsgo/internal/infrastructure/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = repositories.NewGormDatabase()
	e := echo.New()
	e.Use(middleware.CustomLogger)

	log.Fatal(e.Start(":8080"))
}
