package request

import "movie-api/features/transaction"

type ReqTransaction struct {
	AccountID     uint `json:"account_id"`
	PaymentMethod uint `json:"payment_method"`
}

func (reqData *ReqTransaction) ToTransactionCore() transaction.TransactionCore {
	return transaction.TransactionCore{
		AccountID:       reqData.AccountID,
		PaymentMethodID: reqData.PaymentMethod,
	}
}
