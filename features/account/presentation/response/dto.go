package response

import "movie-api/features/account"

type RespAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
