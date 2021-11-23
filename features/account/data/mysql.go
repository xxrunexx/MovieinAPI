package data

import (
	// Import GORM
	"gorm.io/gorm"

	// Import files
	"movie-api/features/account"
)

type AccountData struct {
	db *gorm.DB
}

func NewMySqlAccount(db *gorm.DB) account.Data {
	return &AccountData{db}
}

func (accData *AccountData) InsertAccount(account account.AccountCore) error {
	convData := FromAccountCore(account)

	if err := accData.db.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
