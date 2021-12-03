package request

import "movie-api/features/paymentmethod"

type ReqPaymentmethod struct {
	Name string `json:"name"`
}

func (reqData *ReqPaymentmethod) ToPaymentmethodCore() paymentmethod.PaymentmethodCore {
	return paymentmethod.PaymentmethodCore{
		Name: reqData.Name,
	}
}
