package controllers

import (
	"errors"
	"fmt"
	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentHandlers struct {
	srv service.PaymentService
}

func NewPaymentHandler() *PaymentHandlers {
	gormDb := repositories.NewGormDatabase()
	redisDb := repositories.RedisInit()
	paymentGateway := provider.NewMellatGateway()
	srvPayment := service.NewPaymentService(redisDb, gormDb, paymentGateway)
	return &PaymentHandlers{
		srv: *srvPayment,
	}
}

func RegisterPaymentRoutes(e *echo.Echo) {
	handler := NewPaymentHandler()
	e.GET("/create-payment", handler.CreatePayment)
	e.POST("/verify-payment", handler.VerifyPayment)
}

func (p *PaymentHandlers) CreatePayment(e echo.Context) error {
	err := validators.RequestPaymentValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	reservationId := e.QueryParam("reservation_id")
	userId := e.Get("id").(string)
	redirectLink, err := p.srv.CreateNewPayment(reservationId, userId)
	if err != nil {
		return errors.New("error")
	}
	return e.HTML(http.StatusTemporaryRedirect, redirectLink)

}

func (p *PaymentHandlers) VerifyPayment(e echo.Context) error {
	err := validators.VerifyPaymentValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	refID := e.FormValue("RefId")
	SaleReferenceId := e.FormValue("SaleReferenceId")
	reservationId := e.FormValue("SaleOrderId")
	userId := e.Get("id").(string)
	result, err := p.srv.VerifyPayment(userId, refID, reservationId, SaleReferenceId)
	if err != nil {
		return err
	}

	return e.String(http.StatusOK, fmt.Sprintf("%v", result))
}
