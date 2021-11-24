package business

import "movie-api/features/watchlist"

type WatchlistBusiness struct {
	watchlistData watchlist.Data
}

func NewBusinessWatchlist(watchlistData watchlist.Data) watchlist.Business {
	return &WatchlistBusiness{watchlistData}
}

func (wlBusiness *WatchlistBusiness) CreateWatchlist(wlData watchlist.WatchlistCore) error {
	if err := wlBusiness.watchlistData.InsertWatchlist(wlData); err != nil {
		return err
	}
	return nil
}
