package response

import "movie-api/features/transaction"

type RespTransaction struct {
	ID              uint `json:"id"`
	AccountID       uint `json:"account_id"`
	TotalPrice      int  `json:"price"`
	PaymentMethodID uint `json:"payment_method_id"`
}

func ToTransactionResponse(transaction transaction.TransactionCore) RespTransaction {
	return RespTransaction{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		TotalPrice:      transaction.TotalPrice,
		PaymentMethodID: transaction.PaymentMethodID,
	}
}

func ToTransactionResponseList(trxList []transaction.TransactionCore) []RespTransaction {
	convTrx := []RespTransaction{}

	for _, trx := range trxList {
		convTrx = append(convTrx, ToTransactionResponse(trx))
	}
	return convTrx
}
