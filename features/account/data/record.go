package data

import (
	"movie-api/features/account"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func toAccountRecord(account account.AccountCore) Account {
	return Account{
		Username: account.Username,
		Password: account.Password,
		Email:    account.Email,
	}
}

func toAccountCore(acc Account) account.AccountCore {
	return account.AccountCore{
		Username: acc.Username,
		Password: acc.Password,
		Email:    acc.Email,
	}
}

func toAccountCoreList(accList []Account) []account.AccountCore {
	convAcc := []account.AccountCore{}

	for _, account := range accList {
		convAcc = append(convAcc, toAccountCore(account))
	}
	return convAcc
}
