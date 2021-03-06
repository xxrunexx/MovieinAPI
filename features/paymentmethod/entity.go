package paymentmethod

import "time"

type PaymentmethodCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Untuk layer data
type Data interface {
	InsertPaymentmethod(pmData PaymentmethodCore) (err error)
	SelectPaymentmethod(id int) (PaymentmethodCore, error)
	DeletePaymentmethod(id int) error
}

// Untuk layer business
type Business interface {
	CreatePaymentmethod(pmData PaymentmethodCore) (err error)
	GetPaymentmethod(id int) (PaymentmethodCore, error)
	DeletePaymentmethod(id int) error
}
