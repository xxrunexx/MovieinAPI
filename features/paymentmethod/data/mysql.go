package data

import (
	"movie-api/features/paymentmethod"

	"gorm.io/gorm"
)

type PaymentmethodData struct {
	DB *gorm.DB
}

func NewMySqlPaymentmethod(DB *gorm.DB) paymentmethod.Data {
	return &PaymentmethodData{DB}
}

func (pmData *PaymentmethodData) InsertPaymentmethod(paymentmethod paymentmethod.PaymentmethodCore) error {
	convData := toPaymentmethodRecord(paymentmethod)

	if err := pmData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}

func (pmData *PaymentmethodData) SelectPaymentmethod(id int) (paymentmethod.PaymentmethodCore, error) {
	var singlePaymentMethod Paymentmethod

	err := pmData.DB.Where("id = ?", id).Find(&singlePaymentMethod).Error
	if err != nil {
		return paymentmethod.PaymentmethodCore{}, err
	}
	return toPaymentmethodCore(singlePaymentMethod), nil
}

func (pmData *PaymentmethodData) DeletePaymentmethod(id int) error {
	var singlePaymentmethod Paymentmethod

	err := pmData.DB.Where("id = ?", id).Delete(&singlePaymentmethod).Error
	if err != nil {
		return err
	}
	return nil
}
