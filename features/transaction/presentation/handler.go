package data

import (
	"movie-api/features/transaction"
	"movie-api/features/transaction/presentation/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionBusiness transaction.Business
}

func NewHandlerTransaction(transactionBusiness transaction.Business) *TransactionHandler {
	return &TransactionHandler{transactionBusiness}
}

func (trxHandler *TransactionHandler) CreateTransactionHandler(e echo.Context) error {
	newTransaction := request.ReqTransaction{}

	if err := e.Bind(&newTransaction); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := trxHandler.transactionBusiness.CreateTransaction(newTransaction.ToTransactionCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newTransaction,
	})
}
