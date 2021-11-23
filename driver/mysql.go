package driver

import (
	accdata "movie-api/features/account/data"
	wldata "movie-api/features/watchlist/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:admin@tcp(127.0.0.1)/moviein?parseTime=true"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
	DB.AutoMigrate(&accdata.Account{}, &wldata.Watchlist{})
}
