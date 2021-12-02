package business

import "movie-api/features/paymentmethod"

type PaymentmethodBusiness struct {
	paymentmethodData paymentmethod.Data
}

func NewBusinessPaymentmethod(paymentmethodData paymentmethod.Data) paymentmethod.Business {
	return &PaymentmethodBusiness{paymentmethodData}
}

func (pmBusiness *PaymentmethodBusiness) CreatePaymentmethod(pmData paymentmethod.PaymentmethodCore) error {
	if err := pmBusiness.paymentmethodData.InsertPaymentmethod(pmData); err != nil {
		return err
	}
	return nil
}
