package data

import (
	"movie-api/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	TotalPrice      int `gorm:"default:50000"`
	PaymentMethodID uint
}

func toTransactionRecord(trx transaction.TransactionCore) Transaction {
	return Transaction{
		Model: gorm.Model{
			ID: trx.ID,
		},
		AccountID:       trx.AccountID,
		TotalPrice:      trx.TotalPrice,
		PaymentMethodID: trx.PaymentMethodID,
	}
}

func toTransactionCore(trx Transaction) transaction.TransactionCore {
	return transaction.TransactionCore{
		ID:              trx.ID,
		AccountID:       trx.AccountID,
		PaymentMethodID: trx.PaymentMethodID,
	}
}

func toTransactionCoreList(trxList []Transaction) []transaction.TransactionCore {
	convTrx := []transaction.TransactionCore{}

	for _, transaction := range trxList {
		convTrx = append(convTrx, toTransactionCore(transaction))
	}
	return convTrx
}
