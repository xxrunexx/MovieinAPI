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

func FromWatchlistCore(Data watchlist.WatchlistCore) Watchlist {
	return Watchlist{
		AccountID: Data.AccountID,
		MovieID:   Data.MovieID,
	}
}
