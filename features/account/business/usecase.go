package business

import "movie-api/features/account"

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
