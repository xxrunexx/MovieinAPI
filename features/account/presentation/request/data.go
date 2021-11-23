package request

import (
	"movie-api/features/account"
)

type ReqAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToAccountCore(reqData ReqAccount) account.AccountCore {
	return account.AccountCore{
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
