package business

import (
	"movie-api/features/paymentmethod"
)

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

func (pmBusiness *PaymentmethodBusiness) GetPaymentmethod(id int) (paymentmethod.PaymentmethodCore, error) {
	pmData, err := pmBusiness.paymentmethodData.SelectPaymentmethod(id)
	if err != nil {
		return paymentmethod.PaymentmethodCore{}, err
	}
	return pmData, nil
}

func (pmBusiness *PaymentmethodBusiness) DeletePaymentmethod(id int) error {
	err := pmBusiness.paymentmethodData.DeletePaymentmethod(id)
	if err != nil {
		return err
	}
	return nil
}
