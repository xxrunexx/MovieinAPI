package data

import (
	"movie-api/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	TotalPrice      int
	PaymentMethodID uint
}

type Account struct {
	gorm.Model
	Username string
	Email    string
	// Watchlists []WatchlistCore
}

type PaymentMethod struct {
	gorm.Model
	Name uint
}

// func (acc Account) toTransactionCore() transaction.AccountCore {
// 	return transaction.AccountCore{
// 		ID:       acc.ID,
// 		Username: acc.Username,
// 		Email:    acc.Email,
// 	}
// }

// func (pm PaymentMethod) toTransactionCore() transaction.PaymentMethodCore {
// 	return transaction.PaymentMethodCore{
// 		ID:   pm.ID,
// 		Name: pm.Name,
// 	}
// }

func ToTransactionRecord(transaction transaction.TransactionCore) Transaction {
	return Transaction{
		AccountID:       transaction.AccountID,
		TotalPrice:      transaction.TotalPrice,
		PaymentMethodID: transaction.PaymentMethodID,
	}
}
