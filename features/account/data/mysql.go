package data

import (
	// Import GORM
	"gorm.io/gorm"

	// Import files
	"movie-api/features/account"
)

type AccountData struct {
	DB *gorm.DB
}

func NewMySqlAccount(DB *gorm.DB) account.Data {
	return &AccountData{DB}
}

func (accData *AccountData) InsertAccount(account account.AccountCore) error {
	convData := toAccountRecord(account)

	if err := accData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}

func (accData *AccountData) SelectAccount(account account.AccountCore) ([]account.AccountCore, error) {
	var accounts []Account

	err := accData.DB.Find(&accounts).Error

	if err != nil {
		return nil, err
	}
	return toAccountCoreList(accounts), nil
}
