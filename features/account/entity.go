package account

import "time"

type AccountCore struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type PaymentMethodCore struct {
// 	ID   int
// 	Name string
// }

// Untuk layer data
type Data interface {
	InsertAccount(accData AccountCore) error
	// InsertPayMethod(pmData PaymentMethodCore) error
}

// Untuk layer business
type Business interface {
	CreateAccount(accData AccountCore) error
	// CreatePayMethod(pmData PaymentMethodCore) error
}
