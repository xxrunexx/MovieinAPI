package driver

import (
	accdata "movie-api/features/account/data"
	pmdata "movie-api/features/paymentmethod/data"
	trxdata "movie-api/features/transaction/data"
	wldata "movie-api/features/watchlist/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// For Linux
	// dsn := "root:admin@tcp(127.0.0.1)/moviein?parseTime=true"
	// For Windows
	// dsn := "root:@tcp(127.0.0.1)/moviein?parseTime=true"
	// For RDS
	dsn := "admin:40fied40@tcp(moviein.c4v71mtnu5pg.us-east-2.rds.amazonaws.com)/moviein?parseTime=true"

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
	DB.AutoMigrate(&accdata.Account{}, &wldata.Watchlist{}, &trxdata.Transaction{}, &pmdata.Paymentmethod{})
}
