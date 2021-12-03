package data

import (
	// Import GORM

	"errors"
	"fmt"

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

func (accData *AccountData) SelectAccountByID(id int) (account.AccountCore, error) {
	var singleAccount Account

	err := accData.DB.First(&singleAccount, id).Error

	if singleAccount.Username == "" && singleAccount.ID == 0 {
		return account.AccountCore{}, errors.New("user not found")
	}

	if err != nil {
		return account.AccountCore{}, err
	}

	return toAccountCore(singleAccount), nil
}

func (accData *AccountData) CheckAccount(data account.AccountCore) (account.AccountCore, error) {
	var accountData Account
	err := accData.DB.Where("username = ? and password = ?", data.Username, data.Password).First(&accountData).Error

	// Eliminate null data
	if accountData.Username == "" && accountData.ID == 0 {
		return account.AccountCore{}, errors.New("user not found")
	}

	// Validate with DB
	if err != nil {
		return account.AccountCore{}, err
	}

	return toAccountCore(accountData), nil
}

func (accData *AccountData) UpdateAccount(id int) (account.AccountCore, error) {
	var singleAccount Account
	fmt.Println("Isi single account : ", singleAccount)
	fmt.Println("id : ", id)
	err := accData.DB.Model(&singleAccount).Where("id=?", id).Updates(&singleAccount).Error
	// if singleAccount.Username == "" {
	// 	return account.AccountCore{}, errors.New("user not found")
	// }

	if err != nil {
		return account.AccountCore{}, err
	}

	return toAccountCore(singleAccount), nil
}
