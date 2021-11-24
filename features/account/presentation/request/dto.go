package request

import (
	"movie-api/features/account"
)

type ReqAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AccountAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (data *AccountAuth) ToAccountCore() account.AccountCore {
	return account.AccountCore{
		Username: data.Username,
		Password: data.Password,
	}
}

func (reqData *ReqAccount) ToAccountCore() account.AccountCore {
	return account.AccountCore{
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
