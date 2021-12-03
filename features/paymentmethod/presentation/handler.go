package presentation

import (
	"movie-api/features/paymentmethod"
	"movie-api/features/paymentmethod/presentation/request"
	"movie-api/features/paymentmethod/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentmethodHandler struct {
	paymentmethodBusiness paymentmethod.Business
}

func NewHandlerPaymentmethod(paymentmethodBusiness paymentmethod.Business) *PaymentmethodHandler {
	return &PaymentmethodHandler{paymentmethodBusiness}
}

func (pmHandler *PaymentmethodHandler) CreatePaymentmethodHandler(e echo.Context) error {
	newPaymentmethod := request.ReqPaymentmethod{}

	if err := e.Bind(&newPaymentmethod); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := pmHandler.paymentmethodBusiness.CreatePaymentmethod(newPaymentmethod.ToPaymentmethodCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newPaymentmethod,
	})
}

func (pmHandler *PaymentmethodHandler) GetPaymentmethodHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := pmHandler.paymentmethodBusiness.GetPaymentmethod(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.ToPaymentmethodResponse(data),
	})
}

func (pmHandler *PaymentmethodHandler) DeletePaymentmethodHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi id : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	err = pmHandler.paymentmethodBusiness.DeletePaymentmethod(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "data successfully deleted",
	})
}
