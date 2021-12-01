package response

import "movie-api/features/watchlist"

type RespWatchlist struct {
	ID        int `json:"id"`
	AccountID int `json:"account_id"`
	MovieID   int `json:"movie_id"`
}

func ToWatchlistResponse(watchlist watchlist.WatchlistCore) RespWatchlist {
	return RespWatchlist{
		ID:        watchlist.ID,
		AccountID: watchlist.AccountID,
		MovieID:   watchlist.MovieID,
	}
}

func ToWatchlistResponseList(wlList []watchlist.WatchlistCore) []RespWatchlist {
	convWl := []RespWatchlist{}

	for _, wl := range wlList {
		convWl = append(convWl, ToWatchlistResponse(wl))
	}
	return convWl
}
