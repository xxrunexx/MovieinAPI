package data

import (
	"movie-api/features/transaction"

	"gorm.io/gorm"
)

type TransactionData struct {
	DB *gorm.DB
}

func NewMySqlTransaction(DB *gorm.DB) transaction.Data {
	return &TransactionData{DB}
}

func (trxData *TransactionData) InsertTransaction(transaction transaction.TransactionCore) error {
	convData := toTransactionRecord(transaction)

	if err := trxData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
