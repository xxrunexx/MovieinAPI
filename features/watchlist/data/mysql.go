package data

import (
	"fmt"
	"movie-api/features/watchlist"

	"gorm.io/gorm"
)

type WatchlistData struct {
	DB *gorm.DB
}

func NewMySqlWatchlist(DB *gorm.DB) watchlist.Data {
	return &WatchlistData{DB}
}

func (wlData *WatchlistData) InsertWatchlist(watchlist watchlist.WatchlistCore) error {
	convData := toWatchlistRecord(watchlist)

	if err := wlData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}

func (wlData *WatchlistData) SelectWatchlist(account_id int) ([]watchlist.WatchlistCore, error) {
	var watchlists []Watchlist

	err := wlData.DB.Where("account_id = ?", account_id).Find(&watchlists).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Isi watchlist : ", watchlists)
	return toWatchlistCoreList(watchlists), nil
}

func (wlData *WatchlistData) DeleteWatchlist(id int) error {
	var singleWatchlist Watchlist

	err := wlData.DB.Where("id = ?", id).Delete(&singleWatchlist).Error
	if err != nil {
		return err
	}
	return nil
}
