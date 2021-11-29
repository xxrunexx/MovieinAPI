package data

import (
	"movie-api/features/watchlist"

	"gorm.io/gorm"
)

type Watchlist struct {
	gorm.Model
	AccountID int
	MovieID   int
}

func toWatchlistRecord(watchlist watchlist.WatchlistCore) Watchlist {
	return Watchlist{
		AccountID: watchlist.AccountID,
		MovieID:   watchlist.MovieID,
	}
}

func toWatchlistCore(wl Watchlist) watchlist.WatchlistCore {
	return watchlist.WatchlistCore{
		AccountID: wl.AccountID,
		MovieID:   wl.MovieID,
	}
}

func toWatchlistCoreList(wlList []Watchlist) []watchlist.WatchlistCore {
	convWl := []watchlist.WatchlistCore{}

	for _, watchlist := range wlList {
		convWl = append(convWl, toWatchlistCore(watchlist))
	}
	return convWl
}
