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
