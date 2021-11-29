package response

import "movie-api/features/account"

type RespAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RespAccountLogin struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func ToAccountResponse(account account.AccountCore) RespAccount {
	return RespAccount{
		Username: account.Username,
		Password: account.Password,
		Email:    account.Email,
	}
}

func ToAccountResponseList(accList []account.AccountCore) []RespAccount {
	convAcc := []RespAccount{}

	for _, account := range accList {
		convAcc = append(convAcc, ToAccountResponse(account))
	}
	return convAcc
}

func ToAccountLoginResponse(account account.AccountCore) RespAccountLogin {
	return RespAccountLogin{
		Id:       account.ID,
		Username: account.Username,
		Token:    account.Token,
	}
}
