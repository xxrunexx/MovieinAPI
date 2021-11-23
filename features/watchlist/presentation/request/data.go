package request

import "movie-api/features/watchlist"

type ReqWatchlist struct {
	AccountID int `json:"account_id"`
	MovieID   int `json:"movies_id"`
}

func ToWatchlistCore(reqData ReqWatchlist) watchlist.WatchlistCore {
	return watchlist.WatchlistCore{
		AccountID: reqData.AccountID,
		MovieID:   reqData.MovieID,
	}
}
