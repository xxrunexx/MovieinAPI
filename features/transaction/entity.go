package transaction

import "time"

type TransactionCore struct {
	ID              uint
	AccountID       uint
	TotalPrice      int
	PaymentMethodID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Untuk layer data
type Data interface {
	InsertTransaction(trxData TransactionCore) (err error)
	SelectTransaction(accound_id int) ([]TransactionCore, error)
}

// Untuk layer business
type Business interface {
	CreateTransaction(trxData TransactionCore) (err error)
	GetTransaction(account_id int) ([]TransactionCore, error)
}
