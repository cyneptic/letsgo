package controller

import (
	"errors"
	"fmt"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentHandlers struct {
	srv service.PaymentService
}

func NewPaymentHandler(srv service.PaymentService) *PaymentHandlers {
	return &PaymentHandlers{
		srv: srv,
	}
}

func RegisterPaymentRoutes(e *echo.Echo, srv service.PaymentService) {
	handler := NewPaymentHandler(srv)
	e.GET("/create-payment", handler.CreatePayment)
	e.POST("/verify-payment", handler.VerifyPayment)
}

func (p *PaymentHandlers) CreatePayment(e echo.Context) error {
	reservationId := e.QueryParam("reservation_id")
	if reservationId == "" {
		return echo.ErrBadRequest
	}
	id := e.Get("id").(string)
	redirectLink, err := p.srv.CreateNewPayment(reservationId, id)
	if err != nil {
		return errors.New("error")
	}
	return e.HTML(http.StatusTemporaryRedirect, redirectLink)

}

func (p *PaymentHandlers) VerifyPayment(e echo.Context) error {
	refID := e.FormValue("RefId")
	paymentStatus := e.FormValue("ResCode")
	SaleReferenceId := e.FormValue("SaleReferenceId")
	if paymentStatus != service.SUCCESS_STATUS_CODE {
		return http.ErrAbortHandler
	}
	reservationId := e.FormValue("SaleOrderId")
	id := e.Get("id").(string)
	result, err := p.srv.VerifyPayment(id, refID, reservationId, SaleReferenceId)
	if err != nil {
		return err
	}

	return e.String(http.StatusOK, fmt.Sprintf("%v", result))
}
