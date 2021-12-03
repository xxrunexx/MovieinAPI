package account

import "time"

type AccountCore struct {
	ID        uint
	Username  string
	Password  string
	Email     string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Untuk layer data
type Data interface {
	InsertAccount(accData AccountCore) (err error)
	SelectAccount(AccountCore) (account []AccountCore, err error)
	SelectAccountByID(id int) (account AccountCore, err error)
	CheckAccount(AccountCore) (account AccountCore, err error)
	UpdateAccount(id int) (account AccountCore, err error)
}

// Untuk layer business
type Business interface {
	CreateAccount(accData AccountCore) (err error)
	GetAccount(AccountCore) (account []AccountCore, err error)
	GetAccountByID(id int) (account AccountCore, err error)
	LoginAccount(AccountCore) (account AccountCore, err error)
	UpdateAccount(id int) (account AccountCore, err error)
}
