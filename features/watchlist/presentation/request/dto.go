package request

import "movie-api/features/watchlist"

type ReqWatchlist struct {
	AccountID int `json:"account_id"`
	MovieID   int `json:"movie_id"`
}

func (reqData *ReqWatchlist) ToWatchlistCore() watchlist.WatchlistCore {
	return watchlist.WatchlistCore{
		AccountID: reqData.AccountID,
		MovieID:   reqData.MovieID,
	}
}
