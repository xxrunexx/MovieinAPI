package data

import (
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

	err := wlData.DB.Find(&watchlists, account_id).Error
	if err != nil {
		return nil, err
	}
	return toWatchlistCoreList(watchlists), nil
}
