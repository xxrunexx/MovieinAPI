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
	SelectWatchlist(wlData WatchlistCore) ([]WatchlistCore, error)
}

// Untuk layer business / Service
type Business interface {
	CreateWatchlist(wlData WatchlistCore) error
	GetWatchlist(wlData WatchlistCore) ([]WatchlistCore, error)
}
