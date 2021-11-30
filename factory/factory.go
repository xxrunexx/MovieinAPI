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

	// Tmdb Domain
	tmdbbus "movie-api/features/tmdb/business"
	tmdbdata "movie-api/features/tmdb/data"
	tmdbpres "movie-api/features/tmdb/presentation"
)

type presenter struct {
	AccountPresentation   accpres.AccountHandler
	WatchlistPresentation wlpres.WatchlistHandler
	TmdbPresentation      tmdbpres.TmdbHandler
}

func Init() presenter {
	// Account
	accountData := accdata.NewMySqlAccount(driver.DB)
	accountBusiness := accbus.NewBusinessAccount(accountData)
	// Watchlist
	watchlistData := wldata.NewMySqlWatchlist(driver.DB)
	watchlistBusiness := wlbus.NewBusinessWatchlist(watchlistData)
	// 3rdparty
	tmdbData := tmdbdata.NewData()
	tmdbBusiness := tmdbbus.NewBusinessTmdb(tmdbData)
	return presenter{
		AccountPresentation:   *accpres.NewHandlerAccount(accountBusiness),
		WatchlistPresentation: *wlpres.NewHandlerWatchlist(watchlistBusiness),
		TmdbPresentation:      *tmdbpres.NewHandlerTmdb(tmdbBusiness),
	}
}
