package factory

import (
	// Account Domain
	"movie-api/driver"
	accbus "movie-api/features/account/business"
	accdata "movie-api/features/account/data"
	accpres "movie-api/features/account/presentation"

	// Watchlist Domain
	wlbus "movie-api/features/watchlist/business"
	wldata "movie-api/features/watchlist/data"
	wlpres "movie-api/features/watchlist/presentation"
)

type presenter struct {
	AccountPresentation   accpres.AccountHandler
	WatchlistPresentation wlpres.WatchlistHandler
}

func Init() presenter {
	accountData := accdata.NewMySqlAccount(driver.DB)
	accountBusiness := accbus.NewBusinessAccount(accountData)
	watchlistData := wldata.NewMySqlWatchlist(driver.DB)
	watchlistBusiness := wlbus.NewBusinessWatchlist(watchlistData)
	return presenter{
		AccountPresentation:   *accpres.NewHandlerAccount(accountBusiness),
		WatchlistPresentation: *wlpres.NewHandlerWatchlist(watchlistBusiness),
	}
}
