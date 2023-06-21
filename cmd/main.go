package main

import (
	"log"

	repositories "github.com/cyneptic/letsgo/internal/infrastructure/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = repositories.NewGormDatabase()
	e := echo.New()

	log.Fatal(e.Start(":8080"))
}
