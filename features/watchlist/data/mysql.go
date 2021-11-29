package data

import (
	"movie-api/features/watchlist"

	"gorm.io/gorm"
)

type WatchlistData struct {
	db *gorm.DB
}

func NewMySqlWatchlist(db *gorm.DB) watchlist.Data {
	return &WatchlistData{db}
}

func (wlData *WatchlistData) InsertWatchlist(watchlist watchlist.WatchlistCore) error {
	convData := toWatchlistRecord(watchlist)

	if err := wlData.db.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}

func (wlData *WatchlistData) SelectWatchlist(watchlist watchlist.WatchlistCore) ([]watchlist.WatchlistCore, error) {
	var watchlists []Watchlist

	err := wlData.db.Find(&watchlists).Error
	if err != nil {
		return nil, err
	}
	return toWatchlistCoreList(watchlists), nil
}
