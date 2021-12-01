package data

import (
	"movie-api/features/transaction"
)

type TransactionHandler struct {
	transactionBusiness transaction.Business
}

func NewHandlerTransaction(transactionBusiness transaction.Business) *TransactionHandler {
	return &TransactionHandler{transactionBusiness}
}

// func (trxHandler *TransactionHandler) CreateTransactionHandler(e echo.Context) error {

// }
