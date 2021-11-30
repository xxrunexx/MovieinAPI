package business

import "movie-api/features/transaction"

type TransactionBusiness struct {
	transactionData transaction.Data
}

func NewBusinessTransaction(transactionData transaction.Data) transaction.Business {
	return &TransactionBusiness{transactionData}
}

func (trxBusiness *TransactionBusiness) CreateTransaction(trxData transaction.TransactionCore) error {
	if err := trxBusiness.transactionData.InsertTransaction(trxData); err != nil {
		return err
	}
	return nil
}
