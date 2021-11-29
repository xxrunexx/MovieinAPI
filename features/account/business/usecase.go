package business

import (
	"movie-api/features/account"
	"movie-api/middleware"
)

type AccountBusiness struct {
	accountData account.Data
}

func NewBusinessAccount(accountData account.Data) account.Business {
	return &AccountBusiness{accountData}
}

func (accBusiness *AccountBusiness) CreateAccount(accData account.AccountCore) error {
	if err := accBusiness.accountData.InsertAccount(accData); err != nil {
		return err
	}
	return nil
}

func (accBusiness *AccountBusiness) GetAccount(accData account.AccountCore) ([]account.AccountCore, error) {
	accounts, err := accBusiness.accountData.SelectAccount(accData)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (accBusiness *AccountBusiness) GetAccountByID(id int) (account.AccountCore, error) {
	accountData, err := accBusiness.accountData.SelectAccountByID(id)

	if err != nil {
		return account.AccountCore{}, err
	}
	return accountData, nil
}

func (accBusiness *AccountBusiness) LoginAccount(accData account.AccountCore) (account.AccountCore, error) {
	accountData, err := accBusiness.accountData.CheckAccount(accData)
	if err != nil {
		return account.AccountCore{}, err
	}
	accountData.Token, err = middleware.CreateToken(accData.ID, accData.Username)
	if err != nil {
		return account.AccountCore{}, err
	}
	return accountData, nil
}
