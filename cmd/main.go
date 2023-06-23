package main

import (
	"log"

	"github.com/cyneptic/letsgo/controller"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := repositories.NewPostgres()
	redis := repositories.RedisInit()
	userService := service.NewUserService(db , redis)
	handler := controller.NewHandler(userService, e)
	handler.SetupRoutes()

	log.Fatal(e.Start(":8080"))
}
