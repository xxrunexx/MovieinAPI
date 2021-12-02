package data

import (
	"movie-api/features/transaction"
	"movie-api/features/transaction/presentation/request"
	"movie-api/features/transaction/presentation/response"
	"net/http"
	"strconv"

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

func (trxHandler *TransactionHandler) GetTransactionHandler(e echo.Context) error {
	account_id, err := strconv.Atoi(e.Param("account_id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := trxHandler.transactionBusiness.GetTransaction(account_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.ToTransactionResponseList(data),
	})
}
