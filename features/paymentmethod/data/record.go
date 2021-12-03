package data

import (
	"movie-api/features/paymentmethod"

	"gorm.io/gorm"
)

type Paymentmethod struct {
	gorm.Model
	Name string
}

func toPaymentmethodRecord(pm paymentmethod.PaymentmethodCore) Paymentmethod {
	return Paymentmethod{
		Model: gorm.Model{
			ID: pm.ID,
		},
		Name: pm.Name,
	}
}

func toPaymentmethodCore(pm Paymentmethod) paymentmethod.PaymentmethodCore {
	return paymentmethod.PaymentmethodCore{
		ID:   pm.ID,
		Name: pm.Name,
	}
}

// func toPaymentmethodCoreList(pmList []Paymentmethod) []paymentmethod.PaymentmethodCore {
// 	convPm := []paymentmethod.PaymentmethodCore{}

// 	for _, paymentmethod := range pmList {
// 		convPm = append(convPm, toPaymentmethodCore(paymentmethod))
// 	}
// 	return convPm
// }
