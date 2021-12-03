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

func (trxData *TransactionData) SelectTransaction(accound_id int) ([]transaction.TransactionCore, error) {
	var transactions []Transaction

	err := trxData.DB.Where("account_id = ?", accound_id).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return toTransactionCoreList(transactions), nil
}

func (trxData *TransactionData) DeleteTransaction(id int) error {
	var singleTransaction Transaction

	err := trxData.DB.Where("id = ?", id).Delete(&singleTransaction).Error
	if err != nil {
		return err
	}
	return nil
}
