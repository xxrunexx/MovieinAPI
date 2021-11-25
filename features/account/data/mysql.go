package data

import (
	// Import GORM
	"errors"

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

func (accData *AccountData) SelectAccount(data account.AccountCore) ([]account.AccountCore, error) {
	var accounts []Account

	err := accData.DB.Find(&accounts).Error

	if err != nil {
		return nil, err
	}
	return toAccountCoreList(accounts), nil
}

func (accData *AccountData) CheckAccount(data account.AccountCore) (account.AccountCore, error) {
	var accountData Account

	// Eliminate null data
	if data.Username == "" && data.Password == "" {
		return account.AccountCore{}, errors.New("user not found")
	}

	// Validate with DB
	if err := accData.DB.Where("username = ? and password = ?", data.Username, data.Password).First(&accountData).Error; err != nil {
		return account.AccountCore{}, err
	}
	return toAccountCore(accountData), nil
}
