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
	SelectTransaction() ([]TransactionCore, error)
}

// Untuk layer business
type Business interface {
	CreateTransaction(trxData TransactionCore) (err error)
	GetTransaction() ([]TransactionCore, error)
}
