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

func FromAccountCore(Data account.AccountCore) Account {
	return Account{
		Username: Data.Username,
		Password: Data.Password,
		Email:    Data.Email,
	}
}
