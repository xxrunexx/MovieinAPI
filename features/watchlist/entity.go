package watchlist

import "time"

type WatchlistCore struct {
	ID        int
	AccountID int
	MovieID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Untuk layer data / repo
type Data interface {
	InsertWatchlist(wlData WatchlistCore) (err error)
	SelectWatchlist(account_id int) ([]WatchlistCore, error)
}

// Untuk layer business / Service
type Business interface {
	CreateWatchlist(wlData WatchlistCore) (err error)
	GetWatchlist(account_id int) ([]WatchlistCore, error)
}
