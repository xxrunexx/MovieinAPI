package transaction

import "time"

type TransactionCore struct {
	ID uint
	// Many to One -> account struct
	AccountID  uint
	TotalPrice int
	// Many to One -> PaymentMethod Struct
	PaymentMethodID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type AccountCore struct {
	ID       uint
	Username string
	Email    string
	// Watchlists []WatchlistCore
}

type PaymentMethodCore struct {
	ID   uint
	Name uint
}

// Untuk layer data
type Data interface {
	InsertTransaction(TransactionCore) error
}

// Untuk layer business
type Business interface {
	CreateTransaction(TransactionCore) error
}
