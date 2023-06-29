package validators

import (
	"errors"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RequestPaymentValidator(e echo.Context) error {
	reservationId := e.QueryParam("reservation_id")
	if reservationId == "" {
		return errors.New("reservation_id required")
	}
	_, ok := e.Get("id").(string)
	if !ok {
		return errors.New("user Not found")
	}
	return nil
}

func VerifyPaymentValidator(e echo.Context) error {
	refID := e.FormValue("RefId")
	paymentStatus := e.FormValue("ResCode")
	SaleReferenceId := e.FormValue("SaleReferenceId")
	reservationId := e.FormValue("SaleOrderId")

	if refID == "" || SaleReferenceId == "" || reservationId == "" {
		return errors.New("invalid payment")
	}
	if paymentStatus != entities.SUCCESS_STATUS_CODE {
		return http.ErrAbortHandler
	}
	_, ok := e.Get("id").(string)
	if !ok {
		return errors.New("authentication required")
	}
	return nil
}
