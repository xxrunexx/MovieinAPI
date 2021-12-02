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

func (wlBusiness *WatchlistBusiness) GetWatchlist(account_id int) ([]watchlist.WatchlistCore, error) {
	watchlists, err := wlBusiness.watchlistData.SelectWatchlist(account_id)
	if err != nil {
		return nil, err
	}
	return watchlists, nil
}

func (wlBusiness *WatchlistBusiness) DeleteWatchlist(id int) (watchlist.WatchlistCore, error) {
	watchlistData, err := wlBusiness.watchlistData.DeleteWatchlist(id)
	if err != nil {
		return watchlist.WatchlistCore{}, err
	}
	return watchlistData, nil
}
